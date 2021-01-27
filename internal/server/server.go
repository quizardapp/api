package apiserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/quizardapp/auth-api/internal/store"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var (
	errIncorrectEmailOrPassword = errors.New("Incorrect email or password")
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/register", s.register()).Methods("POST")
	s.router.HandleFunc("/login", s.login()).Methods("POST")
	s.router.HandleFunc("/updatetoken", s.updateAccessToken()).Methods("POST")
	s.router.HandleFunc("/update", s.checkAccessToken(s.updateUser())).Methods("POST")
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) checkAccessToken(next http.HandlerFunc) http.HandlerFunc {

	type request struct {
		ID          string `json:"id"`
		AccessToken string `json:"access_token"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		r.Body.Close() //  must close
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		if err := json.NewDecoder(bytes.NewReader(body)).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		tkn, err := jwt.ParseWithClaims(req.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})

		if err != nil || !tkn.Valid {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})

}

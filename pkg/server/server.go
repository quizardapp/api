package apiserver

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"time"

	"github.com/quizardapp/auth-api/pkg/model"
	"github.com/quizardapp/auth-api/pkg/store"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
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
	s.router.HandleFunc("/register", s.createUser())
}

func (s *server) createUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		uuid, _ := exec.Command("uuidgen").Output()

		u := &model.User{
			ID:           string(uuid),
			Firstname:    req.Firstname,
			Lastname:     req.Lastname,
			Email:        req.Email,
			Password:     req.Password,
			CreationDate: time.Now(),
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}

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

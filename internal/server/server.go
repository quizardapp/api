package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/quizardapp/auth-api/internal/store"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var (
	errorIncorrectEmailOrPassword = errors.New("Incorrect email or password")
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

	userPrefix := "/user"
	s.router.HandleFunc(userPrefix+"/register", s.register()).Methods("POST")
	s.router.HandleFunc(userPrefix+"/login", s.login()).Methods("POST")
	s.router.HandleFunc(userPrefix+"/token", s.updateAccessToken()).Methods("POST")
	s.router.HandleFunc(userPrefix+"/update", s.authenticate(s.updateUser())).Methods("UPDATE")
	s.router.HandleFunc(userPrefix+"/get", s.authenticate(s.getUser())).Methods("POST")

	coursePrefix := "/course"
	s.router.HandleFunc(coursePrefix+"/create", s.authenticate(s.createCourse())).Methods("POST")
	s.router.HandleFunc(coursePrefix+"/get", s.authenticate(s.getCourses())).Methods("POST")
	s.router.HandleFunc(coursePrefix+"/update",s.authenticate(s.updateCourse())).Methods("UPDATE")
	s.router.HandleFunc(coursePrefix+"/delete",s.authenticate(s.deleteCourse())).Methods("DELETE")
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

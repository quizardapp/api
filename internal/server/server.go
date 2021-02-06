package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/quizardapp/auth-api/internal/store"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
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

func (s *server) configureCors() {
	allowedOrigins := handlers.AllowedOrigins([]string{os.Getenv("ALLOWED_ORIGINS")})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodOptions, http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	s.router.Use(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins))
}

func (s *server) configureRouter() {

	s.configureCors()
	s.router.Use(s.setContentType)

	userRouter := s.router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/register", s.register()).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/login", s.login()).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/token", s.updateAccessToken()).Methods("POST")
	userRouter.HandleFunc("/update", s.authenticate(s.updateUser())).Methods("PUT")
	userRouter.HandleFunc("/get", s.authenticate(s.getUser())).Methods("POST")

	courseRouter := s.router.PathPrefix("/course").Subrouter()
	courseRouter.HandleFunc("/create", s.authenticate(s.createCourse())).Methods("POST")
	courseRouter.HandleFunc("/get", s.authenticate(s.getCourses())).Methods("POST")
	courseRouter.HandleFunc("/update", s.authenticate(s.updateCourse())).Methods("PUT")
	courseRouter.HandleFunc("/delete", s.authenticate(s.deleteCourse())).Methods("DELETE")

	moduleRouter := s.router.PathPrefix("/module").Subrouter()
	moduleRouter.HandleFunc("/create", s.authenticate(s.createModule())).Methods("POST")
	moduleRouter.HandleFunc("/get", s.authenticate(s.getModules())).Methods("POST")
	moduleRouter.HandleFunc("/update", s.authenticate(s.updateModule())).Methods("PUT")
	moduleRouter.HandleFunc("/delete", s.authenticate(s.deleteModule())).Methods("DELETE")

	cardRouter := s.router.PathPrefix("/card").Subrouter()
	cardRouter.HandleFunc("/get", s.authenticate(s.getCards())).Methods("POST")
	cardRouter.HandleFunc("/create", s.authenticate(s.createCard())).Methods("POST")
	cardRouter.HandleFunc("/update", s.authenticate(s.updateCard())).Methods("PUT")
	cardRouter.HandleFunc("/delete", s.authenticate(s.deleteCard())).Methods("DELETE")
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

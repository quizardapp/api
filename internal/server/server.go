package server

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inctnce/quizard-api/internal/store"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

// New returns a new server instance
func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store.New(config.Store),
	}
}

// Start runs the server
func (s *Server) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	logrus.Info("server has started on port ", s.config.Port)

	return http.ListenAndServe(":"+s.config.Port, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleBase())
}

func (s *Server) configureStore() error {
	store := store.New(s.config.Store)

	if err := store.Open(); err != nil {
		return err
	}

	s.store = store

	return nil
}

func (s *Server) handleBase() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

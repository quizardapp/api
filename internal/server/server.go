package server

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *Server) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	logrus.Info("server has started on port", s.config.Port)

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

func (s *Server) handleBase() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

package server

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/quizardapp/auth-api/internal/model"
)

func (s *server) createCourse() http.HandlerFunc {

	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
		AccessToken string `json:"access_token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		uuid, err := exec.Command("uuidgen").Output()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		course := &model.Course{
			ID:           strings.TrimSuffix(string(uuid), "\n"),
			Name:         req.Name,
			Description:  req.Description,
			UserID:       req.UserID,
			CreationDate: time.Now(),
		}

		if err := s.store.Course().Create(course); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, course)
	}
}

func (s *server) getCourses() http.HandlerFunc {

	type request struct {
		UserID string `json:"user_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		courses, err := s.store.Course().Read(req.UserID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, courses)

	}
}

func (s *server) updateCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s *server) deleteCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

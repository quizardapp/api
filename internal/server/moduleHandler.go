package server

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/quizardapp/auth-api/internal/model"
)

func (s *server) createModule() http.HandlerFunc {

	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
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

		m := model.Module{
			ID:           strings.TrimSuffix(string(uuid), "\n"),
			Name:         req.Name,
			Description:  req.Description,
			UserID:       req.UserID,
			CreationDate: time.Now(),
		}

		if err := s.store.Module().Create(&m); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, m)
	}
}

func (s *server) getModules() http.HandlerFunc {

	type request struct {
		UserID string `json:"user_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		modules, err := s.store.Module().Read(req.UserID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, modules)
	}
}

func (s *server) deleteModule() http.HandlerFunc {

	type request struct {
		ID string `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Module().Delete(req.ID); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) updateModule() http.HandlerFunc {

	type request struct {
		ModuleID string `json:"id"`
		Value    string `json:"value"`
		Field    string `json:"field"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Module().Update(req.Value, req.Field, req.ModuleID); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

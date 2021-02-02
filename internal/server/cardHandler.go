package server

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/quizardapp/auth-api/internal/model"
)

func (s *server) createCard() http.HandlerFunc {

	type request struct {
		Name     string `json:"name"`
		Content  string `json:"content"`
		ModuleID string `json:"module_id"`
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

		c := &model.Card{
			ID:           strings.TrimSuffix(string(uuid), "\n"),
			Name:         req.Name,
			Content:      req.Content,
			ModuleID:     req.ModuleID,
			CreationDate: time.Now(),
		}

		if err := s.store.Card().Create(c); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusCreated, c)
	}
}

func (s *server) getCards() http.HandlerFunc {

	type request struct {
		ID string `json:"module_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		cards, err := s.store.Card().Read(req.ID)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, cards)
	}
}

func (s *server) updateCard() http.HandlerFunc {

	type request struct {
		CardID string `json:"id"`
		Value  string `json:"value"`
		Field  string `json:"field"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		if err := s.store.Card().Update(req.Value, req.Field, req.CardID); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) deleteCard() http.HandlerFunc {

	type request struct {
		ID string `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		if err := s.store.Card().Delete(req.ID); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)

	}
}

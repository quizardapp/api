package apiserver

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"time"

	"github.com/quizardapp/auth-api/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *server) register() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		uuid, err := exec.Command("uuidgen").Output()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		u := &model.User{
			ID:           string(uuid),
			Firstname:    req.Firstname,
			Lastname:     req.Lastname,
			Email:        req.Email,
			Password:     string(hashed),
			CreationDate: time.Now(),
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}

}

func (s *server) login() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)

		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

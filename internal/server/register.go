package apiserver

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/quizardapp/auth-api/internal/model"
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
			ID:           strings.TrimSuffix(string(uuid), "\n"),
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

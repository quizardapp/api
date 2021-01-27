package apiserver

import (
	"encoding/json"
	"net/http"
)

func (s *server) login() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().Find(req.Email, "email")
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		if err := u.GenerateToken("refresh"); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := u.GenerateToken("access"); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.store.User().Update(u.RefreshToken, "token", u.ID)

		res := &response{u.RefreshToken, u.AccessToken}

		s.respond(w, r, http.StatusOK, res)
	}
}

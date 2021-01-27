package apiserver

import (
	"encoding/json"
	"net/http"
)

func (s *server) updateAccessToken() http.HandlerFunc {
	type request struct {
		ID           string `json:"id"`
		RefreshToken string `json:"token"`
	}

	type response struct {
		AccessToken string
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().Find(req.ID, "id")
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		if err := u.UpdateAccessToken(req.RefreshToken); err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		res := response{u.AccessToken}

		s.respond(w, r, http.StatusOK, res)
	}
}

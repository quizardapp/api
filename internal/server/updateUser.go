package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) updateUser() http.HandlerFunc {

	type request struct {
		ID          string `json:"id"`
		AccessToken string `json:"access_token"`
		Value       string `json:"value"`
		Field       string `json:"field"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.User().Update(req.Value, req.Field, req.ID); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

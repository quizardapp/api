package apiserver

import "net/http"

func (s *server) getUser() http.HandlerFunc {
	type request struct {
		ID          string `json:"id"`
		AccessToken string `json:"access_token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

	}

}

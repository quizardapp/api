package apiserver

import "net/http"

func (s *server) refreshToken() http.HandlerFunc {
	type request struct {
		RefreshToken string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

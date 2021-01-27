package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func (s *server) authenticate(next http.HandlerFunc) http.HandlerFunc {

	type request struct {
		ID          string `json:"id"`
		AccessToken string `json:"access_token"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		req := &request{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(req); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		tkn, err := jwt.ParseWithClaims(req.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})

		if err != nil || !tkn.Valid {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})

}

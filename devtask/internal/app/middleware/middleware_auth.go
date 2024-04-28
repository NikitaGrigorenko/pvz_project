package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func BasicAuth(next http.Handler, expectedUsername string, expectedPassword string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(expectedUsername))
			expectedPasswordHash := sha256.Sum256([]byte(expectedPassword))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, req)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

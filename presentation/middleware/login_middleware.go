package middleware

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/auth"
)

func Login(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetParsedToken(r)

		if err != nil {
			logrus.Errorf("failed to parse token. %v", err)
			http.Error(w, fmt.Sprintf("HTTP Request failed..."), http.StatusInternalServerError)
			return
		}

		if token == nil || !token.Valid {
			logrus.Errorf("token is invalid.")
			http.Error(w, fmt.Sprintf("HTTP Request failed..."), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

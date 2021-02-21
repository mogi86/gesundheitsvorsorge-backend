package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/auth"
)

type HomeController struct{}

type Response struct {
	Message string
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
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

	b, err := json.Marshal(Response{
		Message: "logged in",
	})

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %v", err)
	}
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type HomeController struct{}

type Response struct {
	Message string
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Response{
		Message: "logged in",
	})

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %v", err)
	}
}

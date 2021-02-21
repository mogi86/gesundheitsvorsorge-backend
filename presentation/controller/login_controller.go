package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

type LoginController struct{}

type Jwt struct {
	Token string `json:"token"`
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (l *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: User Login

	key, err := ioutil.ReadFile("./private.key")
	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		logrus.Errorf("failed to parse RSA private key failed. %v\n", err)
		http.Error(w, fmt.Sprintf("HTTP Request failed..."), http.StatusInternalServerError)
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			// TODO: set user_id which is got from datastore
			"user_id": 1,
			// TODO: can be get expire from such as config file
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenStr, err := token.SignedString(parsedKey)
	if err != nil {
		logrus.Errorf("failed to get signed token string. %v\n", err)
		http.Error(w, fmt.Sprintf("HTTP Request failed..."), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	j := &Jwt{
		Token: tokenStr,
	}
	b, err := json.Marshal(j)

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %v\n", err)
	}
}

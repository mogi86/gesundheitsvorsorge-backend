package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"io/ioutil"
	"net/http"
)

func GetParsedToken(r *http.Request) (*jwt.Token, error) {
	key, err := ioutil.ReadFile("./public.key")
	if err != nil {
		logrus.Errorf("cannot read public key. %v\n", err)
		return nil, xerrors.New(fmt.Sprintf("cannot read public key. %v", err))
	}

	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		logrus.Errorf("cannot read parse public key. %v\n", err)
		return nil, xerrors.New(fmt.Sprintf("cannot read parse public key. %v", err))
	}

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, xerrors.New(fmt.Sprintf("method is unexpected. %v", token.Header["alg"]))
		}

		return parsedKey, nil
	})

	return token, nil
}

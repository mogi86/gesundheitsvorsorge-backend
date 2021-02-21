package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mogi86/gesundheitsvorsorge-backend/application/helper"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/request"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/response"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
)

type Jwt struct {
	Token string `json:"token"`
}

type UserController struct {
	usecase usecase.UserInterface
}

func NewUserController(userUseCase usecase.UserInterface) *UserController {
	return &UserController{
		usecase: userUseCase,
	}
}

func (u *UserController) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	logrus.Infof("id: %v", idStr)

	id, _ := strconv.ParseUint(idStr, 10, 64)
	user, err := u.usecase.GetUserById(id)
	if err != nil {
		logrus.Errorf("failed to get user. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	res := &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Mail:      user.Mail,
		Sex:       user.Sex,
		Birthday:  response.CustomDate(user.Birthday),
		Weight:    user.Weight,
		Height:    user.Height,
		Status:    user.Status,
		CreatedAt: response.CustomDateTime(user.CreatedAt),
		UpdatedAt: response.CustomDateTime(user.UpdatedAt),
		TemporaryRegistration: &response.TemporaryRegistration{
			ID:        user.TemporaryRegistration.ID,
			UserID:    user.TemporaryRegistration.UserID,
			Token:     user.TemporaryRegistration.Token,
			ExpireAt:  response.CustomDateTime(user.TemporaryRegistration.ExpireAt),
			CreatedAt: response.CustomDateTime(user.TemporaryRegistration.CreatedAt),
			UpdatedAt: response.CustomDateTime(user.TemporaryRegistration.UpdatedAt),
		},
	}

	b, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %+v\n", err)
	}
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var ruc request.UserCreate

	err := json.NewDecoder(r.Body).Decode(&ruc)
	if err != nil {
		logrus.Errorf("NewDecoder failed. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	birthday, err := time.Parse("2006-01-02", ruc.Birthday)
	if err != nil {
		logrus.Errorf(
			"parse birthday(%v) failed. %v",
			r.PostFormValue("birthday"),
			err,
		)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	weight, err := strconv.ParseFloat(ruc.Weight, 32)
	if err != nil {
		logrus.Errorf("parse weight failed. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	height, err := strconv.ParseFloat(ruc.Height, 32)
	if err != nil {
		logrus.Errorf("parse height failed. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	user := &model.User{
		Password:              helper.ConvertToHash(ruc.Password),
		FirstName:             ruc.FirstName,
		LastName:              ruc.LastName,
		Mail:                  ruc.Mail,
		Sex:                   ruc.Sex,
		Birthday:              birthday,
		Weight:                weight,
		Height:                height,
		Status:                false,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		TemporaryRegistration: model.NewTemporaryRegistration(),
	}

	user, err = u.usecase.CreateUser(user)
	if err != nil {
		logrus.Errorf("parse height failed. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	res := &response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Mail:      user.Mail,
		Sex:       user.Sex,
		Birthday:  response.CustomDate(user.Birthday),
		Weight:    user.Weight,
		Height:    user.Height,
		Status:    user.Status,
		CreatedAt: response.CustomDateTime(user.CreatedAt),
		UpdatedAt: response.CustomDateTime(user.UpdatedAt),
		TemporaryRegistration: &response.TemporaryRegistration{
			ID:        user.TemporaryRegistration.ID,
			UserID:    user.TemporaryRegistration.UserID,
			Token:     user.TemporaryRegistration.Token,
			ExpireAt:  response.CustomDateTime(user.TemporaryRegistration.ExpireAt),
			CreatedAt: response.CustomDateTime(user.TemporaryRegistration.CreatedAt),
			UpdatedAt: response.CustomDateTime(user.TemporaryRegistration.UpdatedAt),
		},
	}

	b, err := json.Marshal(res)
	if err != nil {
		logrus.Errorf("json marshal failed. %v", err)
		http.Error(w, fmt.Sprintf("HTTP Request failed..."), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %v", err)
	}
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var login request.Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		logrus.Errorf("NewDecoder failed. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusBadRequest)
		return
	}

	// user login
	err = u.usecase.Login(login.Mail, login.Password)
	if err != nil {
		logrus.Errorf("failed to login. %v", err)
		http.Error(w, fmt.Sprintf("Bad Request..."), http.StatusUnauthorized)
		return
	}

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

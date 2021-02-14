package controller

import (
	"encoding/json"
	"fmt"
	"github.com/mogi86/gesundheitsvorsorge-backend/application/helper"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/request"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/response"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
)

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
	logrus.Infof("id: %v\n", idStr)

	id, _ := strconv.ParseUint(idStr, 10, 64)
	user := u.usecase.GetUserById(id)
	b, err := json.Marshal(user)

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
		logrus.Errorf("NewDecoder failed. %v\n", err)
		return
	}

	birthday, err := time.Parse("2006-01-02", ruc.Birthday)
	if err != nil {
		logrus.Errorf(
			"parse birthday(%v) failed. %v\n",
			r.PostFormValue("birthday"),
			err,
		)
	}

	weight, err := strconv.ParseFloat(ruc.Weight, 32)
	if err != nil {
		logrus.Errorf("parse weight failed. %v\n", err)
	}

	height, err := strconv.ParseFloat(ruc.Height, 32)
	if err != nil {
		logrus.Errorf("parse height failed. %v\n", err)
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

	user = u.usecase.CreateUser(user)

	res := &response.UserCreate{
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
		logrus.Errorf("json marshal failed. %v\n", err)
		http.Error(w, fmt.Sprintf("HTTP Request failed..."), 500)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %v\n", err)
	}
}

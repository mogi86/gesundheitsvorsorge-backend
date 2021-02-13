package controller

import (
	"encoding/json"
	"fmt"
	"github.com/mogi86/gesundheitsvorsorge-backend/application/helper"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/request"
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
	var ur request.User

	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		logrus.Errorf("NewDecoder failed. %v\n", err)
		return
	}

	birthday, err := time.Parse("2006-01-02", ur.Birthday)
	if err != nil {
		logrus.Errorf(
			"parse birthday(%v) failed. %v\n",
			r.PostFormValue("birthday"),
			err,
		)
	}

	weight, err := strconv.ParseFloat(ur.Weight, 32)
	if err != nil {
		logrus.Errorf("parse weight failed. %v\n", err)
	}

	height, err := strconv.ParseFloat(ur.Height, 32)
	if err != nil {
		logrus.Errorf("parse height failed. %v\n", err)
	}

	user := &model.User{
		Password:  helper.ConvertToHash(ur.Password),
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Mail:      ur.Mail,
		Sex:       ur.Sex,
		Birthday:  birthday,
		Weight:    weight,
		Height:    height,
		Status:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user = u.usecase.CreateUser(user)
	b, err := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, string(b))
	if err != nil {
		logrus.Errorf("return response failed. %v\n", err)
	}
}

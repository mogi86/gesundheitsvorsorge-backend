package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/mogi86/gesundheitsvorsorge-backend/infrastructure/database"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/controller"
)

var (
	db  *sql.DB
	err error
)

func init() {
	user := "gesundheitsvorsorge"
	pass := "gesundheitsvorsorge"
	DBName := "gesundheitsvorsorge_db"
	host := "gesundheitsvorsorge-backend_db_1"
	port := "3306"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, DBName)

	db, err = sql.Open("mysql", dns)
	if err != nil {
		logrus.Errorf("build server failed. %+v\n", err)
	}
}

func main() {
	dbClient := &database.DBClient{
		DB: db,
	}

	// user
	userUseCase := usecase.NewUserUseCase(dbClient)
	userCont := controller.NewUserController(userUseCase)
	http.Handle("/user", userCont)

	// sample
	cont := controller.NewController(&usecase.UseCase{})
	http.Handle("/", cont)

	logrus.Infof("build server...")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		logrus.Errorf("build server failed. %+v\n", err)
	}
}

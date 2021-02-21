package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/middleware"
	"net/http"
	"time"

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
	// TODO: use env values
	user := "gesundheitsvorsorge"
	pass := "gesundheitsvorsorge"
	DBName := "gesundheitsvorsorge_db"
	host := "gesundheitsvorsorge-backend_db_1"
	port := "3306"

	//see: https://github.com/go-sql-driver/mysql#parsetime
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, DBName)

	db, err = sql.Open("mysql", dns)
	if err != nil {
		logrus.Errorf("build server failed. %+v\n", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func main() {
	dbClient := &database.DBClient{
		DB: db,
	}

	// see: https://golang.org/pkg/net/http/#ServeMux
	//   ServeMux is an HTTP request multiplexer
	mux := http.NewServeMux()

	// User
	userUseCase := usecase.NewUserUseCase(dbClient)
	userCont := controller.NewUserController(userUseCase)
	mux.Handle("/user/get", http.HandlerFunc(userCont.FindByID))
	mux.Handle("/user/create", http.HandlerFunc(userCont.Create))
	// Login
	mux.Handle("/login", http.HandlerFunc(userCont.Login))
	// Home
	mux.Handle("/home/index", middleware.Login(
		http.HandlerFunc(controller.NewHomeController().Index)),
	)

	logrus.Infof("build server...")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		logrus.Errorf("build server failed. %+v\n", err)
	}
}

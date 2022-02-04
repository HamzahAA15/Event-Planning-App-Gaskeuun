package main

import (
	"sirclo/config"

	"sirclo/delivery/controllers/graph"
	"sirclo/delivery/router"
	_authRepo "sirclo/repository/auth"
	_commentRepo "sirclo/repository/comment"
	_eventRepo "sirclo/repository/event"
	_participantRepo "sirclo/repository/participant"
	_userRepo "sirclo/repository/user"
	"sirclo/util"

	"github.com/labstack/echo/v4"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	//initiate user model
	// authRepo := auth.New()
	userRepo := _userRepo.New(db)
	authRepo := _authRepo.New(db)
	commentRepo := _commentRepo.New(db)
	participantRepo := _participantRepo.New(db)
	eventRepo := _eventRepo.New(db)
	//create echo http
	e := echo.New()
	client := graph.NewResolver(userRepo, authRepo, commentRepo, participantRepo, eventRepo)
	srv := router.NewGraphQLServer(client)
	router.RegisterPath(e, srv)

	// run server
	e.Logger.Fatal(e.Start(":8080"))
}

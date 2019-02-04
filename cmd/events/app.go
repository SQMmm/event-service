package main

import (
	"github.com/sqmmm/event-service/adapters/handlers/http"
	"github.com/sqmmm/event-service/adapters/repositories/mongoDB"
	"github.com/sqmmm/event-service/usecases/finish"
	"github.com/sqmmm/event-service/usecases/start"
	"gopkg.in/mgo.v2"
)

type app struct {
	session *mgo.Session
	db *mgo.Database
}

func NewApp (session *mgo.Session, dbName string) *app{
	db := session.DB(dbName)
	return &app{session: session, db: db}
}

func (app *app) Start (httpURI string) {
	repo := mongoDB.NewRepository(app.db)
	startUseCase := start.NewUseCase(repo)
	finishUseCase := finish.NewUseCase(repo)

	handler := http.NewHandler(startUseCase, finishUseCase)
	handler.Serve(httpURI)
}

func (app *app) Stop() {
	app.session.Close()
}
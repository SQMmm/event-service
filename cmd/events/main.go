package main

import (
	"github.com/sqmmm/event-service/config"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.Parse()
	if err != nil {
		log.Printf("failed to get config: %s", err)
		os.Exit(1)
	}

	session, err := mgo.Dial(conf.MongoDB.Host+":"+conf.MongoDB.Port)
	if err != nil {
		log.Printf("failed to get mongoDb connection: %s", err)
		os.Exit(1)
	}
	app := NewApp(session, conf.MongoDB.DB)
	
	log.Println("service started")

	go func() {
		app.Start(conf.HTTP.Host + ":" + conf.HTTP.Port)
	}()

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulStop

	app.Stop()
	log.Println("service stopped")
}

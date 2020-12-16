package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/udaya2899/go-gin-starter/configuration"
	"github.com/udaya2899/go-gin-starter/connection"
	"github.com/udaya2899/go-gin-starter/server"
	"github.com/udaya2899/go-gin-starter/storage"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {

	log.Infof("Starting server...")

	if err := run(); err != nil {
		log.Fatalf("Cannot start server, err: %v", err)
	}

}

func run() error {
	config := configuration.New()

	db, err := connection.NewConnection(config.Database)
	if err != nil {
		return err
	}

	repository := storage.New(db)

	go handleShutdown(db)

	s := server.New(repository)

	if err = s.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil {
		return err
	}

	return nil
}

func handleShutdown(db *sql.DB) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	// handle ctrl+c event here
	// for example, close database
	log.Warn("Closing DB connection before complete shutdown")

	if err := db.Close(); err != nil {
		log.Errorf("error while closing the connection to the database: %v", err)
	}

	os.Exit(0)
}

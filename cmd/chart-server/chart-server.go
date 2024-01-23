package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/internal/server"
)

func main() {
	err := database.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Database initialized")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Print("Server started")

	<-done
	log.Print("Server Stopped")
}

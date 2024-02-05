package chartserver

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func RunOrDie() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := serverRun()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Print("Server started")

	<-done
	log.Print("Server Stopped")
}

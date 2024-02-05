package chartserver

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sunglim/chart-server/pkg/cmd/chartserver"
)

func RunOrDie(opts *chartserver.Options) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := serverRun(opts)
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Print("Server started")

	<-done
	log.Print("Server Stopped")
}

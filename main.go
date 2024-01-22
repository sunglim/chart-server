package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sunglim/chart-server/internal"
	"github.com/sunglim/chart-server/internal/options"
)

func main() {
	// write webserver.
	//server.Run()
	//	controller.HandleStockStatistics()

	opts := options.NewOptions()
	cmd := options.InitCommands
	cmd.Run = func(cmd *cobra.Command, args []string) {
		internal.Wrapper(opts)
	}
	opts.AddFlags(cmd)

	if err := opts.Parse(); err != nil {
		fmt.Printf("Failed to parse options: %v\n", err)
		os.Exit(1)
	}
}

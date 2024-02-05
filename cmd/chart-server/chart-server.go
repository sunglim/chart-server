package main

import (
	"github.com/spf13/cobra"
	"github.com/sunglim/chart-server/pkg/cmd/chartserver"
	controller "github.com/sunglim/chart-server/pkg/controller/chartserver"
)

func main() {
	opts := chartserver.NewOptions()
	cmd := chartserver.InitCommands
	cmd.Run = func(cmd *cobra.Command, args []string) {
		controller.RunOrDie(opts)
	}

	opts.AddFlags(cmd)
	opts.Parse()
}

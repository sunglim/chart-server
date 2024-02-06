package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sunglim/chart-server/pkg/cmd/csvinsert"
	controller "github.com/sunglim/chart-server/pkg/controller/csvinsert"
)

func main() {
	opts := csvinsert.NewOptions()
	cmd := csvinsert.InitCommands
	cmd.Run = func(cmd *cobra.Command, args []string) {
		err := controller.NewSvcInsert(opts).Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	opts.AddFlags(cmd)
	err := opts.Parse()
	if err != nil {
		fmt.Println(err)
	}
}

package svcinsert

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "svcinsert",
		Short: "Insert SVC file",
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			fmt.Printf("%s\n", fileName)
			fmt.Println("Not implemented")
		},
	}

	return rootCmd
}

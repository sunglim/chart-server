package options

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Options struct {
	HistoryFiles []string `yaml:"history_files"`

	cmd *cobra.Command
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) AddFlags(cmd *cobra.Command) {
	o.cmd = cmd

	// is called when parsing failed.
	o.cmd.Flags().Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		o.cmd.Flags().PrintDefaults()
	}

	o.cmd.Flags().StringArrayVar(&o.HistoryFiles, "history_files", nil, "history files")
}

func (o *Options) Parse() error {
	err := o.cmd.Execute()
	return err
}

var InitCommands = &cobra.Command{
	Use:   "init",
	Short: "Add arguments for chart server",
	Args:  cobra.NoArgs,
}

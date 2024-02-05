package chartserver

import (
	"github.com/spf13/cobra"
)

// The application. It includes options and command to execute Run().
type Options struct {
	Port int `yaml:"port" json:"port"`

	BaseDirectory string `yaml:"base_directory" json:"base_directory"`

	cmd *cobra.Command
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) AddFlags(cmd *cobra.Command) {
	o.cmd = cmd

	o.cmd.Flags().IntVar(&o.Port, "port", 8080, "Port to listen on")
	o.cmd.Flags().StringVar(&o.BaseDirectory, "base-directory", ".", "Base directory for the server")
}

func (o *Options) Parse() error {
	return o.cmd.Execute()
}

var InitCommands = &cobra.Command{
	Use:   "chartserver",
	Short: "A chart server",
}

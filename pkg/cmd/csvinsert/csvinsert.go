package csvinsert

import (
	"github.com/spf13/cobra"
)

// The application. It includes options and command to execute Run().
type Options struct {
	BaseDirectory string `yaml:"base_directory" json:"base_directory"`
	FileName      string `yaml:"file_name" json:"file_name"`

	cmd *cobra.Command
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) AddFlags(cmd *cobra.Command) {
	o.cmd = cmd

	o.cmd.Flags().StringVar(&o.BaseDirectory, "base-directory", ".", "Base directory for the server")
	o.cmd.Flags().StringVar(&o.FileName, "csv-file-name", "", "CSV file name")
	o.cmd.MarkFlagRequired(o.FileName)
}

func (o *Options) Parse() error {
	return o.cmd.Execute()
}

var InitCommands = &cobra.Command{
	Use:   "chartinsert",
	Short: "A chart data inserter",
}

package csvinsert

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sunglim/chart-server/pkg/controller/csvinsert"
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "csvinsert",
		Short: "Insert SVC file",
		Run: func(cmd *cobra.Command, args []string) {
			fileName := args[0]
			err := NewSvcInsert(fileName).Run()
			if err != nil {
				panic(err)
			}
		},
	}

	return rootCmd
}

type SvcInsert struct {
	fileName string
}

func NewSvcInsert(fileName string) *SvcInsert {
	return &SvcInsert{
		fileName: fileName,
	}
}

func (s *SvcInsert) Run() error {
	err := csvinsert.NewSvcInsert(s.fileName).Run()
	if err != nil {
		fmt.Println("Insert SVC file", "error", err)
	}

	fmt.Println("Insert SVC file")
	return nil
}

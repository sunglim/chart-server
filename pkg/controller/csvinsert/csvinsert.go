package csvinsert

import (
	"fmt"

	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/pkg/cmd/csvinsert"
)

type SvcInsert struct {
	opts *csvinsert.Options
}

func NewSvcInsert(opts *csvinsert.Options) *SvcInsert {
	return &SvcInsert{
		opts: opts,
	}
}

func (s *SvcInsert) Run() error {
	db, err := database.CreateDatabaseFromFilePath(s.opts.BaseDirectory + "/database.db")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	records, err := FromCSVFile(s.opts.BaseDirectory + "/" + s.opts.FileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	database.InsertHistoricalDatas(db, records)

	fmt.Println("records: ", len(records))

	return nil
}

package svcinsert

import (
	"fmt"

	"github.com/sunglim/chart-server/internal/database"
)

type SvcInsert struct {
	fileName string
}

func NewSvcInsert(fileName string) *SvcInsert {
	return &SvcInsert{
		fileName: fileName,
	}
}

func (s *SvcInsert) Run() error {
	db, err := database.CreateDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	records, err := FromCSVFile(s.fileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	database.InsertHistoricalDatas(db, records)

	fmt.Println("records: ", len(records))

	return nil
}

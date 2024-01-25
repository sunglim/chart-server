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
	database.CreateDatabase()

	records, err := FromCSVFile(s.fileName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for record := range records {
		fmt.Println(records[record])
		//	database.Insert(records[record])
	}

	return nil
}

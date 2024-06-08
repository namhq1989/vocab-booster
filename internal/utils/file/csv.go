package appfile

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
)

func ReadCSV[T any](_ *appcontext.AppContext, filePath string) ([]T, error) {
	result := make([]T, 0)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return result, err
	}
	defer func() { _ = file.Close() }()

	err = gocsv.Unmarshal(file, &result)
	return result, err
}

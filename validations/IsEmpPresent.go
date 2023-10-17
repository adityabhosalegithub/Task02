package validations

import (
	"encoding/csv"
	"os"
	"strconv"
)

func IsEmpPresent(id int) (bool, error) {
	CSVFile := "data/emp.csv"
	file, err := os.Open(CSVFile)
	if err != nil {
		return false, CustomError{Message: "Error in opening csvfile"}
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return false, CustomError{Message: "Error in reading csvfile"}
	}
	for _, record := range records {
		if len(record) > 0 {
			recordID, _ := strconv.Atoi(record[0])
			if recordID == id {
				return true, nil
			}
		}
	}
	return false, CustomError{Message: "Employee record with given id not found"}
}

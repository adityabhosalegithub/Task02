package csvutil

import (
	"encoding/csv"
	"main/searchutil"
	"main/validations"
	"main/view"
	"os"
)

func DeleteEmployeeByID(employeeID int) error {

	// CSVFile := "data/emp.csv"
	file, err := os.OpenFile(view.CSVFile, os.O_RDWR, 0644)
	//os.O_RDWR allows file to opened for reading and writing, and 0644 sets the file permissions to allow the owner to read and write
	if err != nil {
		return validations.CustomError{Message: "CSV File Not Opening"}
	}
	defer file.Close()

	reader := csv.NewReader(file) //a new CSV reader that reads from the provided file
	records, _ := reader.ReadAll()

	pos, err := searchutil.GetRecordPosition(records, employeeID)
	if err != nil {
		return err
	}
	// by append deleting the record at position
	records = append(records[:pos], records[pos+1:]...)

	err = file.Truncate(0)
	if err != nil {
		return validations.CustomError{Message: "Error truncating file"}
	}
	file.Seek(0, 0) // set the file pointer to the beginning

	writer := csv.NewWriter(file)
	for _, r := range records {
		err := writer.Write(r)
		if err != nil {
			return validations.CustomError{Message: "Error In Writer"}
		}
	}
	writer.Flush()

	return nil
}

package view

import (
	"encoding/csv"
	"os"
)

const CSVFile = "data/emp.csv"

func init() {
	//CSVFile := "data/emp.csv"
	err := CheckOrCreateCSVFile(CSVFile)
	if err != nil {
		panic("Error in Loding CSV File") //special function used for package init and cannot return any values
	}
}

func CheckOrCreateCSVFile(filePath string) error {
	_, err := os.Stat(filePath) //to check file is present or not

	if os.IsNotExist(err) { // returns bool on directory does not exist.
		//supports err return by os pkg
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		header := []string{"ID", "FirstName", "LastName", "Email", "Password", "PhoneNo", "Role", "Salary", "Birthday"} // adding emp header
		writer.Write(header)
	}
	return nil
}

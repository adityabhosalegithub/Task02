package csvutil

import (
	"encoding/csv"
	"main/view"
	"os"
	"strconv"
	"time"
)

func LoadCSVData() error { //func used in printing all data
	file, err := os.Open(view.CSVFile)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file) //a new CSV reader that reads from the provided file
	records, err := reader.ReadAll()
	//reader.ReadAll()  reads all the records from the CSV file and returns them as  two-dimensional slice of strings [][]string
	/*
			Name, Age, City ,
			adi, 20, satara
			it look like
			[][]string{
		        []string{"Name", "Age", "City"},
		        []string{"adi", "21", "satara"},
			}

	*/
	if err != nil {
		return err
	}

	for _, record := range records {
		salary, _ := strconv.ParseFloat(record[7], 64)
		birthday, _ := time.Parse("2006-01-02", record[8])
		id, _ := strconv.Atoi(record[0])
		employee := view.Employee{
			ID:        id,
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
			Password:  record[4],
			PhoneNo:   record[5],
			Role:      record[6],
			Salary:    salary,
			Birthday:  birthday,
		}

		view.Emp = append(view.Emp, employee)
	}
	view.Emp = view.Emp[1:] //to remove header like id,name,emil

	return nil
}

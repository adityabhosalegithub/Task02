package csvutil

import (
	"encoding/csv"
	"main/view"
	"os"
	"strconv"
)

// saveCSVData saves data to a CSV file.
func AddDataInCsv(newEmployee view.Employee) {
	// Open the CSV file in append mode
	file, err := os.OpenFile(view.CSVFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//os.O_APPEND is flag used to open the file in append mode.When a file is opened with this flag, any writes to the file will be appended to the end rather than overwriting existing content.
	//os.O_WRONLY flag used to open the file in write-only mode, allowing you to write data to the file.
	//os.ModeAppend is file mode that indicates that file should be opened in append mode.
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file) //ceate a CSV writer
	defer writer.Flush()

	record := []string{ // convert the new employeedata to strings
		strconv.Itoa(newEmployee.ID),
		newEmployee.FirstName,
		newEmployee.LastName,
		newEmployee.Email,
		newEmployee.Password,
		newEmployee.PhoneNo,
		newEmployee.Role,
		strconv.FormatFloat(newEmployee.Salary, 'f', 2, 64),
		newEmployee.Birthday.Format("2006-01-02"),
	}

	err = writer.Write(record) // write the new record to the CSV file
	if err != nil {
		panic(err)
	}

}

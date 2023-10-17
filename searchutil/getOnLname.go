package searchutil

import (
	"main/utility"
	"main/validations"
	"main/view"
	"strings"
)

func GetOnLastname(lastname string, records [][]string) ([]view.Employee, error) {
	lastname = strings.ToLower(lastname)
	var matchingEmployees []view.Employee

	if len(lastname) > 0 {
		for _, record := range records {
			field := strings.ToLower(record[2]) // LastName is at index 2
			if field == lastname {
				empl := utility.StringToEmployee(record)
				matchingEmployees = append(matchingEmployees, empl)
			}
		}
	}

	if len(matchingEmployees) == 0 {
		return []view.Employee{}, validations.CustomError{Message: "Employee record with given LastName not found"}
	}

	return matchingEmployees, nil
}

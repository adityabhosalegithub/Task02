package searchutil

import (
	"main/utility"
	"main/validations"
	"main/view"
	"strings"
)

func GetOnFirstname(firstname string, records [][]string) ([]view.Employee, error) {
	firstname = strings.ToLower(firstname)
	var matchingEmployees []view.Employee

	if len(firstname) > 0 {
		for _, record := range records {
			field := strings.ToLower(record[1]) // FirstName is at index 1
			if field == firstname {
				empl := utility.StringToEmployee(record)
				matchingEmployees = append(matchingEmployees, empl)
			}
		}
	}

	if len(matchingEmployees) == 0 {
		return []view.Employee{}, validations.CustomError{Message: "Employee record with given FirstName not found"}
	}

	return matchingEmployees, nil
}

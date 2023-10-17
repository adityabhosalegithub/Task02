package searchutil

import (
	"main/utility"
	"main/validations"
	"main/view"
	"strings"
)

func GetOnEmail(email string, records [][]string) (view.Employee, error) {
	email = strings.ToLower(email)
	if len(email) > 0 {
		for _, record := range records {

			field := strings.ToLower(record[3]) // Email is at index 3
			if field == email {
				empl := utility.StringToEmployee(record)
				return empl, nil
			}
		}
	}
	return view.Employee{}, validations.CustomError{Message: "Employee record with given Email not found"}
}

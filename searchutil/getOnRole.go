package searchutil

import (
	"main/utility"
	"main/validations"
	"main/view"
	"strings"
)

func GetOnRole(role string, records [][]string) ([]view.Employee, error) {
	p := false
	role = strings.ToLower(role)
	if len(role) > 0 {
		for _, record := range records {

			field := strings.ToLower(record[6]) // role is at index 6
			if field == role {
				empl := utility.StringToEmployee(record)
				view.Emp = append(view.Emp, empl)
				p = true
			}
		}
	}
	if len(view.Emp) > 1 {
		view.Emp = view.Emp[1:]
	}
	if !p {
		return []view.Employee{}, validations.CustomError{Message: "Employee record with given Role not found"}
	}

	return view.Emp, nil
}

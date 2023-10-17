package searchutil

import (
	"main/utility"
	"main/validations"
	"main/view"
	"strconv"
)

func GetOnID(id int, records [][]string) (view.Employee, error) {
	if id == 0 {
		return view.Employee{}, validations.CustomError{Message: "Employee ID Can Not Be Zero"}
	} else {
		for _, record := range records {
			field, _ := strconv.Atoi(record[0]) // Email is at index 3
			if field == id {
				empl := utility.StringToEmployee(record)
				return empl, nil
			}
		}
	}
	return view.Employee{}, validations.CustomError{Message: "Employee record with given ID not found"}
}

package utility

import (
	"main/view"
	"strconv"
	"time"
)

// convert the slice of strings to employee struct fields
func StringToEmployee(fields []string) view.Employee {
	id, _ := strconv.Atoi(fields[0])
	salary, _ := strconv.ParseFloat(fields[7], 64)
	birthday, _ := time.Parse("2006-01-02", fields[8])

	return view.Employee{
		ID:        id,
		FirstName: fields[1],
		LastName:  fields[2],
		Email:     fields[3],
		Password:  fields[4],
		PhoneNo:   fields[5],
		Role:      fields[6],
		Salary:    salary,
		Birthday:  birthday,
	}
}

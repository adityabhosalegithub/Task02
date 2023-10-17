package utility

import (
	"main/view"
	"strconv"
)

// need method to covert data sutatabile to csv
// convert the employee struct fields to a formatted string
func EmployeeToStringArr(employee view.Employee) []string {
	return []string{
		strconv.Itoa(employee.ID),
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.Password,
		employee.PhoneNo,
		employee.Role,
		strconv.FormatFloat(employee.Salary, 'f', 2, 64),
		employee.Birthday.Format("2006-01-02"),
	}
}

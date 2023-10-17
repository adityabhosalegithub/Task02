package validations

import (
	"fmt"
	"main/view"
	"net/http"
	"reflect"
	"strings"
)

// to check all updating fields
func ValidateAllUpdatingDatafield(reqid int, w http.ResponseWriter, newEmployee view.Employee) bool {
	id := newEmployee.ID
	fmt.Println("Received ID:", id)
	if err := IsEmptyInt(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	if err := IsIdISUnique(id); err != nil && reqid != id { //emp can keep same id so we take requie id
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	firstname := newEmployee.FirstName
	firstname = strings.TrimSpace(firstname)
	fmt.Println("Received First Name:", firstname)
	if err := IsValidFirstName(firstname); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		//http.Error is a function in net/http package .it sends an HTTP response with a specified status code and body.
		// err.Error() take error message from the error object
		return false
	}
	lastname := newEmployee.LastName
	lastname = strings.TrimSpace(lastname)
	fmt.Println("Received Last Name:", lastname)
	if err := IsValidLastName(lastname); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	email := newEmployee.Email
	email = strings.TrimSpace(email)
	fmt.Println("Received Email:", email)
	if err := IsValidEmail(email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	// if err := IsUniqueEmail(email); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return false
	// }
	existingEmail, ok := CheckCurrentMail(id)
	if !ok {
		// id not found, no need to check the current email
		// check if the new email is unique
		err := IsUniqueEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return false
		}
	} else {
		if existingEmail != email {
			err := IsUniqueEmail(email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return false
			}
		}
	}

	password := newEmployee.Password
	password = strings.TrimSpace(password)
	fmt.Println("Received Password:", password)
	if err := IsValidPassword(password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	phoneno := newEmployee.PhoneNo
	phoneno = strings.TrimSpace(phoneno)
	fmt.Println("Received PhoneNo:", phoneno)
	if _, err := IsValidPhoneno(phoneno); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	role := newEmployee.Role
	role = strings.TrimSpace(role)
	fmt.Println("Received Role:", role)
	if err := IsValidRole(role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	salary := newEmployee.Salary
	fmt.Println(reflect.TypeOf(salary))
	fmt.Println("Received salary:", salary)
	if err := IsValidSalary(salary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	date := newEmployee.Birthday
	//dateString = strings.TrimSpace(dateString)//cannot use strings.TrimSpace(dateString) (value of type string) as time.Time value in assignment
	// date, err := time.Parse("2006-01-02", dateString)
	fmt.Println("Received Birthday:", date)
	if _, err := CalculateAge(date); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	fmt.Println("Employee Details Enterd SucessFullly")

	return true
}

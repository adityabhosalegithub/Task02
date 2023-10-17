package curd

import (
	"encoding/json"
	"fmt"
	"main/csvutil"
	"main/validations"
	"main/view"
	"net/http"
	"reflect"
	"strings"
)

func AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter-- w is interface provided by the net/http package that allows u to construct an HTTP response
	//r *http.Request ---r is a pointer to an http.Request struct, which represents the incoming HTTP request.
	//r - this struct contains information about the request, such as the method, headers, URL, and request body.
	var newEmployee view.Employee
	err := json.NewDecoder(r.Body).Decode(&newEmployee)
	//json.NewDecoder(r.Body)--creates new JSON decoder that reads from the request's body.
	//.Decode(&newEmployee): decodes JSON data from the request body into the newEmployee variable and & we passing pointer to newEmployee
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := newEmployee.ID
	fmt.Println("Received ID:", id)
	if err := validations.IsEmptyInt(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validations.IsIdISUnique(id); err != nil {
		http.Error(w, err.Error(), http.StatusConflict) //409 conflict
		return
	}

	firstname := newEmployee.FirstName
	firstname = strings.TrimSpace(firstname)
	fmt.Println("Received First Name:", firstname)
	if err := validations.IsValidFirstName(firstname); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //400 badrequest
		//http.Error is a function in net/http package .it sends an HTTP response with a specified status code and body.
		// err.Error() take error message from the error object
		return
		// no return statement here, so the code will continue to execute
		//from down code he will be executed regardless of whether there was an error or not
	}

	lastname := newEmployee.LastName
	lastname = strings.TrimSpace(lastname)
	fmt.Println("Received Last Name:", lastname)
	if err := validations.IsValidLastName(lastname); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := newEmployee.Email
	email = strings.TrimSpace(email)
	fmt.Println("Received Email:", email)
	if err := validations.IsValidEmail(email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validations.IsUniqueEmail(email); err != nil {
		http.Error(w, err.Error(), http.StatusConflict) //409
		return
	}

	password := newEmployee.Password
	password = strings.TrimSpace(password)
	fmt.Println("Received Password:", password)
	if err := validations.IsValidPassword(password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	phoneno := newEmployee.PhoneNo
	phoneno = strings.TrimSpace(phoneno)
	fmt.Println("Received PhoneNo:", phoneno)
	if _, err := validations.IsValidPhoneno(phoneno); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	role := newEmployee.Role
	role = strings.TrimSpace(role)
	fmt.Println("Received Role:", role)
	if err := validations.IsValidRole(role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	salary := newEmployee.Salary
	fmt.Println(reflect.TypeOf(salary))
	fmt.Println("Received salary:", salary)
	if err := validations.IsValidSalary(salary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	date := newEmployee.Birthday
	//dateString = strings.TrimSpace(dateString)//cannot use strings.TrimSpace(dateString) (value of type string) as time.Time value in assignment
	// date, err := time.Parse("2006-01-02", dateString)
	fmt.Println("Received Birthday:", date)
	if _, err := validations.CalculateAge(date); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //ttp.Error method sends a plain text response by default
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee added successfully!")

	csvutil.AddDataInCsv(newEmployee)
}

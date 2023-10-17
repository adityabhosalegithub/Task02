package validations

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// all cutom validations
func IsValidFirstName(firstname string) error {
	if firstname != "" {
		return nil
	}
	return CustomError{Message: "Firstname cannot be empty. Enter Name again"}
	// createing CustomError instance with a specific error message and return
}

func IsValidLastName(lastname string) error {
	if lastname != "" {
		return nil
	}
	return CustomError{Message: "Lastname cannot be empty. Enter Name again"}
}

func IsValidEmail(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return nil
	}
	return CustomError{Message: "Email should contain @ and . Enter valid email"}
}

func IsValidPassword(password string) error {
	if password != "" {
		return nil
	}
	return CustomError{Message: "Password cannot be empty. Enter password again"}
}

func IsValidPhoneno(phoneno string) (int, error) {
	if len(phoneno) != 10 {
		return 0, CustomError{Message: "phone number must be 10 digits long"}
	}
	phone, err := strconv.ParseInt(phoneno, 10, 64)
	if err != nil {
		return 0, CustomError{Message: "failed to convert phone number to integer"}
	}
	return int(phone), nil
}

func IsValidRole(role string) error {
	if role == "" {
		return CustomError{Message: "Role cannot be empty. Enter role again"}
	}
	role = strings.ToLower(role)
	if role == "admin" || role == "manager" || role == "developer" || role == "tester" {
		return nil
	} else {
		return CustomError{Message: "failed given role can't be accepted"}
	}

}

func IsValidSalary(salary float64) error {
	if salary == 0 {
		return CustomError{Message: "Salary cannot be ZERO. Enter salary again"}
	}
	// convertedSalary, err := strconv.ParseFloat(salary, 64)
	return nil
}

func ParseDate(dateString string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return time.Time{}, CustomError{Message: "Entered Date not in proper format. Please enter the date again."}
	}
	return date, nil
}

func CalculateAge(date time.Time) (int, error) {
	currentTime := time.Now()
	age := currentTime.Year() - date.Year()
	if currentTime.Before(time.Date(currentTime.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)) {
		age-- // to handle cases where the birthday has not occurred yet in the current year.
	}

	if age >= 18 {
		return age, nil
	} else {
		return age, CustomError{Message: "Employee is not 18 years old yet. Please enter a valid birthdate."}
	}
}

func CheckNilErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func IsValidInt(input string) (int, error) {
	num, err := strconv.ParseInt(input, 10, 64) //converting string to int
	if err != nil {
		return 0, CustomError{Message: "failed to convert phone number to integer"}
	}
	return int(num), nil
}

func IsEmptyString(input string) error {
	str := strings.TrimSpace(input)
	if str == "" {
		return CustomError{Message: "Empty value entered...Enter value again"}
	}
	return nil
}

func IsUniqueEmail(email string) error { //to check email is unique
	CSVFile := "data/emp.csv"
	file, err := os.Open(CSVFile)
	if err != nil {
		return CustomError{Message: "Error in opening csvfile"}
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	email = strings.ToLower(email)
	if len(email) > 0 {
		for _, record := range records {
			field := strings.ToLower(record[3]) // Email is at index 3
			if field == email {
				return CustomError{Message: "Enter Unique Email"}
			}
		}
	}
	return nil
}

func IsEmptyInt(input int) error {
	if input == 0 {
		return CustomError{Message: "Empty value entered...Enter value again"}
	}
	return nil
}
func IsIdISUnique(id int) error {
	p, _ := IsEmpPresent(id)
	if p {
		return CustomError{Message: "Id Already In Use ...Provide Unique id"}
	}
	return nil
}

func CheckCurrentMail(id int) (string, bool) {
	CSVFile := "data/emp.csv"
	file, _ := os.Open(CSVFile)
	defer file.Close()
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	for _, record := range records {
		if record[0] == strconv.Itoa(id) {
			return record[3], true
		}
	}
	return "", false
}

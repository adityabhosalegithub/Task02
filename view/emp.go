package view

import (
	"time"
)

type Employee struct {
	ID        int       `json:"ID"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	PhoneNo   string    `json:"PhoneNo"`
	Role      string    `json:"Role"`
	Salary    float64   `json:"Salary"`
	Birthday  time.Time `json:"Birthday"`
}

var Emp []Employee // a global var of type slice of Employee

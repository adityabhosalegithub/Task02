package curd

import (
	"encoding/csv"
	"encoding/json"
	"main/searchutil"
	"main/validations"
	"net/http"
	"os"
	"strings"
)

func SearchByFnameHandler(w http.ResponseWriter, firstName string) {
	firstName = strings.TrimSpace(firstName)
	if err := validations.IsValidFirstName(firstName); err != nil {
		http.Error(w, validations.CustomError{Message: "In URL Firstname cannot be empty. Send Correct Request Again"}.Error(), http.StatusBadRequest)
		return
	}

	CSVFile := "data/emp.csv"
	file, err := os.OpenFile(CSVFile, os.O_RDWR, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	emp, err := searchutil.GetOnFirstname(firstName, records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

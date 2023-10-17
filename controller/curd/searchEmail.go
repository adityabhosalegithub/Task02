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

func SearchByEmailHandler(w http.ResponseWriter, email string) {
	email = strings.TrimSpace(email)
	if err := validations.IsValidEmail(email); err != nil {
		http.Error(w, validations.CustomError{Message: "In URL Email cannot be empty. Send Correct Request Again"}.Error(), http.StatusBadRequest)
		return
	}

	CSVFile := "data/emp.csv"
	file, err := os.Open(CSVFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	emp, err := searchutil.GetOnEmail(email, records)
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

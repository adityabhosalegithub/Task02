package curd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"main/searchutil"
	"main/validations"
	"net/http"
	"os"
)

func SearchByIdHandler(w http.ResponseWriter, input string) {
	// empid := r.URL.Query().Get("id")
	empid := input
	id, err := validations.IsValidInt(empid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Received ID:", id)

	CSVFile := "data/emp.csv"
	file, err := os.Open(CSVFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file) //a new CSV reader that reads from the provided file
	records, _ := reader.ReadAll()

	emp, err := searchutil.GetOnID(id, records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// marshal the data to JSON
	data, err := json.Marshal(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the response headers
	w.Header().Set("Content-Type", "application/json")
	// send the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(data) //body
}

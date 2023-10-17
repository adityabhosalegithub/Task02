package curd

import (
	"encoding/json"
	"main/csvutil"
	"main/view"
	"net/http"
)

func PrintEmployeeDataHandler(w http.ResponseWriter, r *http.Request) {
	err := csvutil.LoadCSVData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //500
		return
	}
	// mmarshal the data to JSON
	data, err := json.Marshal(view.Emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// set the response headers
	w.Header().Set("Content-Type", "application/json")
	// send the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	view.Emp = view.Emp[:0] //making empty

}

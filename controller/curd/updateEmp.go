package curd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"main/searchutil"
	"main/utility"
	"main/validations"
	"main/view"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// http://localhost3000/updateEmployee/1
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	//taking id fron url
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	fmt.Println("p0=", parts[0])
	fmt.Println("p1=", parts[1])
	fmt.Println("p2=", parts[2])

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID in URL", http.StatusBadRequest)
		return
	}
	fmt.Println("id to be Updated", id)

	CSVFile := "data/emp.csv"
	//os.Open(CSVfile) giving error err-write data/emp.csv: Access is denied.
	file, err := os.OpenFile(CSVFile, os.O_RDWR, 0644)
	//os.O_RDWR allows file to opened for reading and writing, and 0644 sets the file permissions to allow the owner to read and write
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	reader := csv.NewReader(file) //a new CSV reader that reads from the provided file
	records, _ := reader.ReadAll()

	pos, err := searchutil.GetRecordPosition(records, id) //passing[][]csv and id
	//in pos thre is record id
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var newEmployee view.Employee
	err = json.NewDecoder(r.Body).Decode(&newEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if ok := validations.ValidateAllUpdatingDatafield(id, w, newEmployee); !ok { //passing writer and emp
		// emp can keep same id so we pass id whuch was in requestheader
		return
	}
	fmt.Println("record at pos is goint to be changed", pos)
	// override the value of the second record
	record := utility.EmployeeToStringArr(newEmployee)
	records[pos] = record
	fmt.Println("changed record", records[pos])

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}
	file.Seek(0, 0) // set the file pointer to the beginning

	writer := csv.NewWriter(file)
	for _, r := range records {
		err := writer.Write(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	writer.Flush()

	// if err := writer.Error(); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee Updated successfully!")

}

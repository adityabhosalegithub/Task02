package curd

import (
	"fmt"
	"main/csvutil"
	"net/http"
	"strconv"
	"strings"
)

func DeleteEmployeeDataHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	fmt.Println("p2=", parts[2])
	employeeID, _ := strconv.Atoi(parts[2])
	if employeeID == 0 {
		http.Error(w, "Error Employee Id Empty", http.StatusInternalServerError)
		return
	}
	err := csvutil.DeleteEmployeeByID(employeeID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting employee: %s", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // if deletion was successful, send a success message
	fmt.Fprintf(w, "Employee with ID %d has been deleted successfully", employeeID)
}

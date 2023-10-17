package controller

import (
	"main/controller/curd"
	"net/http"

	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()

	// Route for /Employee
	router.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			curd.PrintEmployeeDataHandler(w, r)
		} else if r.Method == http.MethodPost {
			curd.AddEmployeeHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route for /Employee/{id}
	router.HandleFunc("/employee/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			curd.UpdateEmployeeHandler(w, r)
		case http.MethodDelete:
			curd.DeleteEmployeeDataHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route for /Employee/search
	router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		queryValues := r.URL.Query()
		if len(queryValues) == 0 {
			http.Error(w, "No search parameters provided", http.StatusBadRequest)
			return
		}
		if len(queryValues) != 1 {
			http.Error(w, "Please provide exactly one search parameter", http.StatusBadRequest)
			return
		}
		for key, value := range queryValues {
			switch key {
			case "id":
				curd.SearchByIdHandler(w, value[0])
			case "firstName":
				curd.SearchByFnameHandler(w, value[0])
			case "lastName":
				curd.SearchByLastnameHandler(w, value[0])
			case "email":
				curd.SearchByEmailHandler(w, value[0])
			case "role":
				curd.SearchByRoleHandler(w, value[0])
			default:
				http.Error(w, "Invalid search parameter", http.StatusBadRequest)
				return
			}
		}
	})

	return router

}

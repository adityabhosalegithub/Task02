package searchutil

import (
	"main/validations"
	"strconv"
)

func GetRecordPosition(records [][]string, id int) (int, error) {
	// printRecords(records)
	if len(records) > 0 {
		for i, record := range records {
			// fmt.Printf("Record %d:\n", i)
			for _, field := range record {
				// fmt.Printf("\tField %d: %s\n", j, field)
				if rid, _ := strconv.Atoi(field); id == rid {
					return i, nil // Return the position (record index)
				}
				// break
			}
		}
	}
	return 0, validations.CustomError{Message: "Employee record with given id not found"}
}

/*
func printRecords(records [][]string) {
	for i, record := range records {
		fmt.Printf("Record %d:\n", i)
		for j, field := range record {
			fmt.Printf("\tField %d: %s\n", j, field)
		}
	}
} */

package dataset

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// MovieVector is a map of movie IDs to a slice of tags
var MovieVector map[string][]string

func CleanDataset() {

	csvfile, err := os.Open("dataset/tags.csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)

	// Read all the records
	records, err := r.ReadAll()
	if err != nil {
		log.Println(err)
		return
	}

	MovieVector := make(map[string][]string)

	// Map the movie ID to the tags
	for _, record := range records {
		if !StringInSlice(record[2], MovieVector[record[1]]) {
			MovieVector[record[1]] = append(MovieVector[record[1]], record[2])
		}
	}

	for k, v := range MovieVector {
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}

}

// StringInSlice function to check if a string is in a slice of strings
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

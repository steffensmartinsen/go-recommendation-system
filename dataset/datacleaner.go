package dataset

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
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
		MovieVector[record[1]] = append(MovieVector[record[1]], record[2])
	}

	// Print the 10 first movie IDs and tags
	for i := 0; i < 10; i++ {
		log.Println(MovieVector[strconv.Itoa(i)])
	}

}

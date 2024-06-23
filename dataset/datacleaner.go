package dataset

import (
	"encoding/csv"
	"log"
	"os"
)

func GenerateTagVector() []string {

	csvfile, err := os.Open("dataset/tags.csv")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)

	// Read all the records
	records, err := r.ReadAll()
	if err != nil {
		log.Println(err)
		return nil
	}

	// Create a map to store the movie ID and the tags
	MovieVector := make(map[string][]string)

	// Map the movie ID to the tags
	for _, record := range records {
		if !StringInSlice(record[2], MovieVector[record[1]]) {
			MovieVector[record[1]] = append(MovieVector[record[1]], record[2])
		}
	}

	tagVector := []string{}
	for _, tags := range MovieVector {
		for _, tag := range tags {
			if !StringInSlice(tag, tagVector) {
				tagVector = append(tagVector, tag)
			}
		}
	}

	return tagVector
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

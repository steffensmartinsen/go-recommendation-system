package dataset

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

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

	fmt.Println(records[0])
	fmt.Println(records[1])

}

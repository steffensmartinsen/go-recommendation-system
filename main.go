package main

import (
	"go-recommendation-system/dataset"
	"log"
	"time"
)

func main() {

	start := time.Now()
	dataset.CleanDataset()

	elapsed := time.Since(start)
	log.Printf("Data cleaning took %s", elapsed)

}

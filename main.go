package main

import (
	"go-recommendation-system/dataset"
	"log"
	"time"
)

func main() {

	start := time.Now()
	tags := dataset.GenerateTagVector()

	log.Printf(tags[0])

	elapsed := time.Since(start)
	log.Printf("Data cleaning took %s", elapsed)

}

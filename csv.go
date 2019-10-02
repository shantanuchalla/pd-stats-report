package main

import (
	"encoding/csv"
	"log"
	"os"
)

func writeToCsv(values [][]string) {
	file, err := os.Create("stats.csv")
	if err != nil {
		log.Fatalln(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range values {
		err := writer.Write(value)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var title string

func main() {
	// Open the file
	csvfile, err := os.Open("sample.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	mandlMap := make(map[string]int)
	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		// if err != nil {
		// 	log.Fatal(err)
		// }

		if strings.Contains(strings.ToLower(record[0]), strings.ToLower("Log Book Not Uploaded")) {
			title = record[0]
		}

		if len(record) > 2 {
			if strings.ToLower(record[1]) == "sangareddy" {
				mandlMap[record[5]] = mandlMap[record[5]] + 1
			}
		}
	}

	file, _ := os.Create("result.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{title})
	writer.Write([]string{"Mandal", "Habitation count"})
	count := 0
	for k, v := range mandlMap {
		writer.Write([]string{k, fmt.Sprintf("%v", v)})
		count = count + v
	}
	writer.Write([]string{"Total", fmt.Sprintf("%v", count)})
}

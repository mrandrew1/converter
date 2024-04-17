package createWordFiles

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

const prePath = "./uploads/data/"

func getDataFromCsv(path string) ([][]string, []string) {
	f, err := os.Open(prePath + path)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	csvDataArray, columns := createCSVDataArray(data)
	return csvDataArray, columns
}

func createCSVDataArray(data [][]string) ([][]string, []string) {
	var csvDataArray [][]string
	var columns []string

	for i, line := range data {
		if i == 0 {
			for _, field := range line {
				columns = append(columns, strings.Trim(field, " "))
			}
		}
		if i > 0 { // omit header line
			csvDataArray = append(csvDataArray, line)
		}
	}

	return csvDataArray, columns
}

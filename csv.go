package main

import (
	"bufio"
	"encoding/csv"
	"strconv"
)

func parseCsv(r *bufio.Reader) (Stats, error) {
	csvReader := csv.NewReader(r)

	// Read the header
	header, err := csvReader.Read()
	if err != nil {
		panic(err)
	}

	var s Stats

	numberOfColumns := len(header)

	// Read the data
	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		if len(record) != numberOfColumns {
			panic("Invalid CSV file")
		}

		for _, value := range record {
			// Check if the value is numeric, value is a string
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				s.NumStats.Sum += f
				s.NumStats.Count++
			} else {
				s.StrStats.CharCount += len(value)
				s.StrStats.WordCount++
			}

		}

		s.NumStats.Mean = s.NumStats.Sum / float64(s.NumStats.Count)
	}

	return s, nil
}

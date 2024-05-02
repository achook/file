package main

import (
	"bufio"
	"encoding/csv"
)

func parseCsv(r *bufio.Reader) (Stats, error) {
	csvReader := csv.NewReader(r)

	// Read the header
	header, err := csvReader.Read()
	if err != nil {
		panic(err)
	}

	rs := NewRawStrStats()
	rn := RawNumStats{}

	numberOfColumns := len(header)

	// Read the data
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return Stats{}, err
			}
		}

		if len(record) != numberOfColumns {
			panic("Invalid CSV file")
		}

		for _, value := range record {
			// Check if the value is numeric, value is a string
			handleTextWord(value, &rs, &rn)

		}
	}

	return NewStats(&rs, &rn), nil
}

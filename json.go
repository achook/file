package main

import (
	"bufio"
	"encoding/json"
	"io"
)

func traverse(v interface{}, stats *Stats) *Stats {
	switch vv := v.(type) {
	case map[string]interface{}:
		for _, value := range vv {
			// Check if the value is numeric
			if num, ok := value.(float64); ok {
				handleNumber(num, stats)
			} else if _, ok := value.(bool); ok {
				// TODO: handle boolean
			} else if str, ok := value.(string); ok {
				handleString(str, stats)
			} else {
				// If not numeric, boolean, or string, recursively traverse the value
				traverse(value, stats)
			}
		}

	case []interface{}:
		// Traverse through each element of the array recursively
		for _, value := range vv {
			traverse(value, stats)
		}
	}

	return stats
}

func parseJSON(r *bufio.Reader) (Stats, error) {

	// Unmarshal JSON data into a generic interface{}
	var jsonData interface{}
	decoder := json.NewDecoder(&io.LimitedReader{N: 20, R: r})
	err := decoder.Decode(&jsonData)
	if err != nil {
		return Stats{}, err
	}

	s := Stats{}
	traverse(jsonData, &s)

	return s, nil
}

package main

import (
	"bufio"
	"encoding/json"
)

func traverse(v interface{}, rs *RawStrStats, rn *RawNumStats) {
	switch vv := v.(type) {
	case map[string]interface{}:
		for _, value := range vv {
			// Check if the value is numeric
			if num, ok := value.(float64); ok {
				handleNumberField(num, rn)
			} else if str, ok := value.(string); ok {
				handleTextField(str, rs)
			} else {
				// If not numeric, boolean, or string, recursively traverse the value
				traverse(value, rs, rn)
			}
		}

	case []interface{}:
		// Traverse through each element of the array recursively
		for _, value := range vv {
			traverse(value, rs, rn)
		}
	}
}

func parseJSON(r *bufio.Reader) (Stats, error) {

	// Unmarshal JSON data into a generic interface{}
	var jsonData interface{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&jsonData)
	if err != nil {
		return Stats{}, err
	}

	rs := NewRawStrStats()
	rn := RawNumStats{}

	// Traverse the JSON data
	traverse(jsonData, &rs, &rn)

	return NewStats(&rs, &rn), nil
}

package main

import (
	"bufio"
)

func parseText(r *bufio.Reader) (Stats, error) {
	rs := NewRawStrStats()
	rn := RawNumStats{}
	// Read the data

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return Stats{}, err
			}
		}

		handleTextLine(line, &rs, &rn)
	}

	return NewStats(&rs, &rn), nil
}

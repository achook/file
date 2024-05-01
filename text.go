package main

import (
	"bufio"
	"strconv"
	"strings"
)

func parseText(r *bufio.Reader) (Stats, error) {
	s := Stats{}
	// Read the data
	var lineLengths []int
	var wordLengths []int

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		// Get line length
		ll := len(line)
		lineLengths = append(lineLengths, ll)
		s.StrStats.LineCount++
		s.StrStats.CharCount += ll

		words := strings.Split(line, " ")
		for _, word := range words {
			wl := len(word)

			if wl > 0 {
				wordLengths = append(wordLengths, wl)
				s.StrStats.WordCount++
			}

			// Try to parse the word as a number
			if val, err := strconv.ParseFloat(word, 64); err == nil {
				// If successful, increment the number of words
				s.NumStats.Count++
				s.NumStats.Sum += val
			}

		}
	}

	// Calculate mean
	s.NumStats.Mean = s.NumStats.Sum / float64(s.NumStats.Count)

	// Calculate average line length, average word length, mode line length, mode word length
	s.StrStats.MeanLineLength = float64(s.StrStats.CharCount) / float64(s.StrStats.LineCount)
	s.StrStats.MeanWordLength = float64(s.StrStats.CharCount) / float64(s.StrStats.WordCount)

	// Calculate mode line length
	lineLengthMap := make(map[int]int)
	for _, ll := range lineLengths {
		lineLengthMap[ll]++
	}
	modeLineLength := 0
	maxLineLength := 0
	for ll, count := range lineLengthMap {
		if count > maxLineLength {
			maxLineLength = count
			modeLineLength = ll
		}
	}

	// Calculate mode word length
	wordLengthMap := make(map[int]int)
	for _, wl := range wordLengths {
		wordLengthMap[wl]++
	}
	modeWordLength := 0
	maxWordLength := 0
	for wl, count := range wordLengthMap {
		if count > maxWordLength {
			maxWordLength = count
			modeWordLength = wl
		}
	}

	s.StrStats.ModeWordLength = float64(modeWordLength)
	s.StrStats.ModeLineLength = float64(modeLineLength)

	return s, nil
}

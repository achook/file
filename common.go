package main

import (
	"math"
	"strconv"
	"strings"
)

func handleTextField(v string, rs *RawStrStats) {
	wl := len(v)
	if val, ok := rs.wordLengthMap[wl]; ok {
		rs.wordLengthMap[wl] = val + 1
	} else {
		rs.wordLengthMap[wl] = 1
	}

	// Check if the word is an email
	rs.emailCount += countEmails(v)
}

func handleTextLine(v string, rs *RawStrStats, rn *RawNumStats) {
	// Get line length
	ll := len(v)
	if val, ok := rs.lineLengthMap[ll]; ok {
		rs.lineLengthMap[ll] = val + 1
	} else {
		rs.lineLengthMap[ll] = 1
	}

	rs.lineCount++

	// Get words
	words := splitLine(v)

	for _, word := range words {
		handleTextWord(word, rs, rn)
	}
}

func handleTextWord(v string, rs *RawStrStats, rn *RawNumStats) {
	wl := len(v)
	if val, ok := rs.wordLengthMap[wl]; ok {
		rs.wordLengthMap[wl] = val + 1
	} else {
		rs.wordLengthMap[wl] = 1
	}

	// Check if the word is an email
	rs.emailCount += countEmails(v)

	// Try to parse the word as a number
	if num, err := strconv.ParseFloat(v, 64); err == nil {
		rn.numbers = append(rn.numbers, num)
	}
}

func handleNumberField(v float64, rn *RawNumStats) {
	rn.numbers = append(rn.numbers, v)
}

func splitLine(v string) []string {
	return strings.Split(v, " ")
}

// calculateStdDev calculates the standard deviation of a list of integers
func calculateStdDev(l []int) float64 {
	sum := 0
	for _, v := range l {
		sum += v
	}

	mean := float64(sum) / float64(len(l))

	variance := 0.0
	for _, v := range l {
		variance += (float64(v) - mean) * (float64(v) - mean)
	}

	r := math.Sqrt(variance / float64(len(l)))

	if math.IsNaN(r) {
		return 0
	}

	return r
}

// calculateFloat64StdDev calculates the standard deviation of a list of float64
func calculateFloat64StdDev(l []float64) float64 {
	sum := 0.0
	for _, v := range l {
		sum += v
	}

	mean := sum / float64(len(l))

	variance := 0.0
	for _, v := range l {
		variance += (v - mean) * (v - mean)
	}

	r := math.Sqrt(variance / float64(len(l)))

	if math.IsNaN(r) {
		return 0
	}

	return r
}

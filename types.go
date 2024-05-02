package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"math"
)

type Stats struct {
	NumStats NumStats `json:"num_stats"`
	StrStats StrStats `json:"str_stats"`
}

type NumStats struct {
	Sum    float64 `json:"sum"`
	Count  int     `json:"count"`
	Mean   float64 `json:"mean"`
	Mode   float64 `json:"mode"`
	StdDev float64 `json:"standard_deviation"`
}

type StrStats struct {
	WordCount        int     `json:"word_count"`
	CharCount        int     `json:"char_count"`
	LineCount        int     `json:"line_count"`
	MeanWordLength   float64 `json:"mean_word_length"`
	MeanLineLength   float64 `json:"mean_line_length"`
	ModeWordLength   float64 `json:"mode_word_length"`
	ModeLineLength   float64 `json:"mode_line_length"`
	StdDevLineLength float64 `json:"std_dev_line_length"`
	StdDevWordLength float64 `json:"std_dev_word_length"`
}

type RawStrStats struct {
	wordLengthMap map[int]int
	lineLengthMap map[int]int
	lineCount     int
	emailCount    int
}

func NewRawStrStats() RawStrStats {
	return RawStrStats{
		wordLengthMap: make(map[int]int),
		lineLengthMap: make(map[int]int),
	}
}

type RawNumStats struct {
	numbers []float64
}

// NewStats returns a new Stats object made from the given RawStrStats and RawNumStats
func NewStats(rs *RawStrStats, rn *RawNumStats) Stats {
	s := Stats{}
	ss := StrStats{}
	ns := NumStats{}

	// Calculate StrStats
	ss.WordCount = len(rs.wordLengthMap)
	ss.LineCount = rs.lineCount

	// Calculate char count
	for k, v := range rs.wordLengthMap {
		ss.CharCount += k * v
	}

	// Calculate mean word length
	ss.MeanWordLength = float64(ss.CharCount) / float64(ss.WordCount)
	if math.IsNaN(ss.MeanWordLength) {
		ss.MeanWordLength = 0

	}

	// Calculate mean line length
	for k, v := range rs.lineLengthMap {
		ss.MeanLineLength += float64(k) * float64(v)
	}
	ss.MeanLineLength /= float64(ss.LineCount)
	if math.IsNaN(ss.MeanLineLength) {
		ss.MeanLineLength = 0
	}

	// Calculate mode word length
	maxVal := 0
	for k, v := range rs.wordLengthMap {
		if v > maxVal {
			maxVal = v
			ss.ModeWordLength = float64(k)
		}
	}

	// Calculate mode line length
	maxVal = 0
	for k, v := range rs.lineLengthMap {
		if v > maxVal {
			maxVal = v
			ss.ModeLineLength = float64(k)
		}
	}

	// Calculate standard deviation of line length
	var lineLengths []int
	for k, v := range rs.lineLengthMap {
		for i := 0; i < v; i++ {
			lineLengths = append(lineLengths, k)
		}
	}
	ss.StdDevLineLength = calculateStdDev(lineLengths)

	// Calculate standard deviation of word length
	var wordLengths []int
	for k, v := range rs.wordLengthMap {
		for i := 0; i < v; i++ {
			wordLengths = append(wordLengths, k)
		}
	}
	ss.StdDevWordLength = calculateStdDev(wordLengths)

	// Calculate NumStats
	ns.Count = len(rn.numbers)
	for _, v := range rn.numbers {
		ns.Sum += v
	}

	ns.Mean = ns.Sum / float64(ns.Count)
	if math.IsNaN(ns.Mean) {
		ns.Mean = 0
	}

	// Calculate mode
	numMap := make(map[float64]int)
	maxVal = 0
	for _, v := range rn.numbers {
		numMap[v]++
		if numMap[v] > maxVal {
			maxVal = numMap[v]
			ns.Mode = v
		}
	}

	// Calculate standard deviation
	ns.StdDev = calculateFloat64StdDev(rn.numbers)

	s.NumStats = ns
	s.StrStats = ss

	return s
}

type Metrics struct {
	csvFilesCount  *prometheus.CounterVec
	jsonFilesCount *prometheus.CounterVec
	textFilesCount *prometheus.CounterVec

	responseTime *prometheus.HistogramVec
	fileSize     *prometheus.HistogramVec
}

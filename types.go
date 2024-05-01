package main

import "github.com/prometheus/client_golang/prometheus"

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

type RawStats struct {
	wordLengthMap map[int]int
	lineLengthMap map[int]int
}

type metrics struct {
	csvFilesCount  *prometheus.CounterVec
	jsonFilesCount *prometheus.CounterVec
	textFilesCount *prometheus.CounterVec

	responseTime *prometheus.HistogramVec
	fileSize     *prometheus.HistogramVec
}

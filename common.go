package main

import "strings"

func handleString(v string, stats *Stats) *Stats {
	// split into words
	words := strings.Split(v, " ")
	for _, word := range words {

		wl := len(word)

		if wl > 0 {
			stats.StrStats.WordCount++
			stats.StrStats.CharCount += wl
		}
	}

	return stats
}

func handleNumber(v float64, stats *Stats) *Stats {
	stats.NumStats.Sum += v
	stats.NumStats.Count++
	stats.NumStats.Mean = stats.NumStats.Sum / float64(stats.NumStats.Count)

	return stats

}

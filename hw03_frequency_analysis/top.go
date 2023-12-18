package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	words := strings.Fields(input)
	wordCounts := countWords(words)

	wordsList := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		wordsList = append(wordsList, word)
	}

	sort.Slice(wordsList, func(i, j int) bool {
		return wordCounts[wordsList[i]] > wordCounts[wordsList[j]] ||
			(wordCounts[wordsList[i]] == wordCounts[wordsList[j]] && wordsList[i] < wordsList[j])
	})
	
	resultLength := 10
	if len(wordsList) < 10 {
		resultLength = len(wordsList) 
	}

	result := make([]string, resultLength)
	for i := range result {
		result[i] = wordsList[i]
	}

	return result
}

func countWords(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

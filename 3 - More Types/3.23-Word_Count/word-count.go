package main

import (
	// "golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	word_count := make(map[string]int)

	for _, value := range words {
		word_count[value] += 1
	}

	return word_count
}

func main() {
	// wc.Test(WordCount)
}

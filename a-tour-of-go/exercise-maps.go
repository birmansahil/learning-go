package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	wordCounter := make(map[string]int)

	words := strings.Fields(s)

	for _, v := range words {
		wordCounter[v]++
	}

	return wordCounter
}

func main() {
	fmt.Println(wordCount("I am learning learning learning Go"))
}

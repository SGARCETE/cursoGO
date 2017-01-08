package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	m := make(map[string]int)

	for i := range words {
		m[words[i]] = count(words[i], words)
	}
	return m
}

func count(s string, words []string) int {
	contador := 0
	for i := range words {
		if s == words[i] {
			contador++
		}
	}
	return contador
}

func main() {
	wc.Test(WordCount)
}
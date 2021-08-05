package main

import (
	"fmt" 
	"strings"
	"os"
)

func main() {
	src:= readInput()
	wordcounter := make(map[string]int)
	words := strings.Fields(src)
	for _, word := range words {
		if _, ok := wordcounter[word]; !ok {
			wordcounter[word] = 0
		}
		wordcounter[word] += 1
	}
	for key, val := range wordcounter {
		fmt.Println(key, ":", val)
	}
}

func readInput() string {
	src := os.Args[1]
	return src
}
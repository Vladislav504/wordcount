package main

import (
	"fmt" 
	"strings"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	wordcounter := make(map[string]int)
	words := strings.Fields(text)
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
package main

import (
	"fmt" 
	"strings"
	"os"
)

func main() {
	src:= readInput()
	words := strings.Fields(src)
	fmt.Println(len(words))

}

func readInput() string {
	src := os.Args[1]
	return src
}
package main

import (
	"log"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
}

func main(){

	log.Print(WordCount("aij jhjk aij"))

}

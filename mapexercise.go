package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	myMap := make(map[string]int)

	inputStrings := strings.Split(s, " ")
	fmt.Println(inputStrings)

	for _, v := range inputStrings {
		value, isPresent := myMap[v]
		if isPresent {
			myMap[v] = value + 1
		} else {
			myMap[v] = 1
		}
	}
	return myMap
}

func main() {
	wc.Test(WordCount)
}

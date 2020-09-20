package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	cnts := make(map[string]int)
	for _, f := range strings.Fields(s) {
		if _, ok := cnts[f]; ok {
			cnts[f]++
		} else {
			cnts[f] = 1
		}
	}
	return cnts
}

func main() {
	wc.Test(WordCount)
}

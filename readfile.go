package main

import (
	"fmt"
	"os"
)

func fileRead() []string {
	var countPrefix int = 0
	var txt_words []string
	wordMore := ""
	for {
		_, err := fmt.Fscan(os.Stdin, &wordMore)
		if err != nil {
			break
		}
		txt_words = append(txt_words, wordMore)
		countPrefix++
		if countPrefix == *prefix {
			break
		}
	}
	return txt_words
}

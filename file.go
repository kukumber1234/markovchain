package main

import (
	"fmt"
	"os"
)

func fullFile() []string {
	var txt_words []string
	wordMore := ""
	for {
		_, err := fmt.Fscan(os.Stdin, &wordMore)
		if err != nil {
			break
		}
		txt_words = append(txt_words, wordMore)
	}
	return txt_words
}

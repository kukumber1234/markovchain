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
	if len(txt_words) == 1 || len(txt_words) == 0 {
		fmt.Fprintln(os.Stderr, "Do not use echo please")
		os.Exit(1)
	}
	return txt_words
}

func fileRead(save_txt []string) []string {
	var txt_words_part []string
	for i := 0; i < *prefix; i++ {
		txt_words_part = append(txt_words_part, save_txt[i])
	}
	return txt_words_part
}

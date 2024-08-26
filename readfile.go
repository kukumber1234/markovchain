package main

import (
	"bufio"
	"os"
)

func fileRead(fileName string) []string {
	var countPrefix int = 0
	f, _ := os.Open(fileName)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	savePrefixWords := []string{}
	for s.Scan() {
		word := s.Text()
		savePrefixWords = append(savePrefixWords, word)
		countPrefix++
		if countPrefix == *prefix {
			break
		}
	}
	return savePrefixWords
}

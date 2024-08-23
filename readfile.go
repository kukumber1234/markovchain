package main

import (
	"bufio"
	"os"
)

func fileRead(a string) []string {
	var countPrefix int = 0
	f, _ := os.Open(a)
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	d := []string{}
	for s.Scan() {
		word := s.Text()
		d = append(d, word)
		countPrefix++
		if countPrefix == *prefix {
			break
		}
	}
	return d
}

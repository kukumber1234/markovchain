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
	d := []string{}
	for s.Scan() {
		d = append(d, s.Text())
		if countPrefix == *prefix {
			break
		}
		countPrefix++
	}
	return d
}

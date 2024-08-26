package main

import (
	"fmt"
	"os"
	"strings"
)

func checkNum(defaultNumber *int) {
	if *defaultNumber < 0 {
		fmt.Fprintln(os.Stderr, "Given number can't be negative")
		os.Exit(1)
	}
	if *defaultNumber > 10000 {
		fmt.Fprintln(os.Stderr, "Given number can't be more than 10,000")
		os.Exit(1)
	}
	if *defaultNumber == 0 {
		fmt.Fprintln(os.Stderr, "Given number can't be equal to 0")
		os.Exit(1)
	}
}

func checkPref(prefix *int) {
	if *prefix < 0 || *prefix > 5 {
		fmt.Fprintln(os.Stderr, "Prefix number must be between 0 <= x <= 5")
		os.Exit(1)
	}
}

func checkWord(save_txt, fileSave []string) {
	prefixWordString := strings.Fields(*prefixWord)
	if len(prefixWordString) > 1 && len(prefixWordString) < 5 {
		*prefix = len(prefixWordString)
		save_txt = fullFile()
		for i := 0; i < len(prefixWordString); i++ {
			fileSave = append(fileSave, prefixWordString[i])
		}
		mapSave(save_txt, fileSave)
	} else {
		fmt.Fprintln(os.Stderr, "Prefix word len must be between 2 <= x <= 5")
		os.Exit(1)
	}
}

func checkLen() {
	prefixWordString := strings.Fields(*prefixWord)
	if len(prefixWordString) != *prefix {
		fmt.Fprintln(os.Stderr, "Len of prefix word and prefix must be equal")
		os.Exit(1)
	}
}

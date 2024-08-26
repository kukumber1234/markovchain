package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func mapSave(save_txt []string, fileSave []string) {
	wordMap := make(map[string][]string)
	emptySlice := make([]string, *prefix)
	var randWordChoose int

	for k := 0; k < len(save_txt); k++ {

		key := strings.Join(emptySlice, " ")
		wordMap[key] = append(wordMap[key], save_txt[k])
		newWord := save_txt[k]
		if len(emptySlice) > 0 {
			emptySlice = append(emptySlice[1:], newWord)
		}
	}

	if *defaultNumber < len(fileSave) {
		for i := 0; i < *defaultNumber; i++ {
			fmt.Print(fileSave[i] + " ")
		}
	} else {
		for i := 0; i < len(fileSave); i++ {
			fmt.Print(fileSave[i] + " ")
		}
	}

	for i := 0; i < *defaultNumber-len(fileSave); i++ {
		arrayToString := strings.Join(fileSave, " ")
		if len(wordMap[arrayToString]) > 0 {
			randWordChoose = rand.Intn(len(wordMap[arrayToString]))
			nextWord := wordMap[arrayToString][randWordChoose]
			fmt.Print(nextWord + " ")
			fileSave = append(fileSave[1:], nextWord)
		} else {
			fmt.Println(":this word do not exist or do not have continue in the text")
			os.Exit(1)
		}
	}
	fmt.Println()
}

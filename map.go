package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func mapSave(save_txt []string, fileSave []string) {
	wordMap := make(map[string][]string)
	// var check_identically_word bool
	// var save_txt_part []string
	emptySlice := make([]string, *prefix)
	var randWordChoose int

	for k := 0; k < len(save_txt); k++ {

		key := strings.Join(emptySlice, " ")
		wordMap[key] = append(wordMap[key], save_txt[k])
		newWord := save_txt[k]
		// fmt.Println(wordMap[arrayToString][0])
		if len(emptySlice) > 0 {
			emptySlice = append(emptySlice[1:], newWord)
		} else {
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

	// fmt.Println(wordMap)
	for i := 0; i < *defaultNumber-len(fileSave); i++ {
		arrayToString := strings.Join(fileSave, " ")
		if len(wordMap[arrayToString]) > 0 {
			// fmt.Println(len(wordMap[arrayToString]))
			randWordChoose = rand.Intn(len(wordMap[arrayToString]))
			nextWord := wordMap[arrayToString][randWordChoose]
			fmt.Print(nextWord + " ")
			// fmt.Print(wordMap[arrayToString][randWordChoose] + " ")
			fileSave = append(fileSave[1:], nextWord)
		} else {
			fmt.Println()
			fmt.Println("It is the very last word in the text")
			os.Exit(0)
		}
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func mapSave(save_txt []string, fileSave []string) {
	wordMap := make(map[string][]string)
	var check_identically_word bool
	var save_txt_part []string
	var randWordChoose int

	if *defaultNumber < len(fileSave) {
		for i := 0; i < *defaultNumber; i++ {
			fmt.Print(fileSave[i] + " ")
		}
	} else {
		for i := 0; i < len(fileSave); i++ {
			fmt.Print(fileSave[i] + " ")
		}
	}

	for k := len(fileSave); k < *defaultNumber; k++ {
		arrayToString := strings.Join(fileSave, " ")

		for m := 0; m < len(save_txt); m++ {
			save_txt_part = append(save_txt_part, save_txt[m])
			if len(save_txt_part) == *prefix {
				check_identically_word = true
				for j := 0; j < *prefix; j++ {
					if save_txt_part[j] != fileSave[j] {
						check_identically_word = false
						break
					}
				}
				if check_identically_word && m+1 < len(save_txt) {
					wordMap[arrayToString] = append(wordMap[arrayToString], save_txt[m+1])
				}
				save_txt_part = save_txt_part[1:]
			}

		}
		if len(wordMap[arrayToString]) > 0 {
			randWordChoose = rand.Intn(len(wordMap[arrayToString]))
			fmt.Print(wordMap[arrayToString][randWordChoose] + " ")
			fileSave = append(fileSave[1:], wordMap[arrayToString][randWordChoose])
			// fmt.Println(fileSave)
		} else {
			break
		}
	}
	fmt.Println()
}

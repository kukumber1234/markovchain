package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func mapSave(save_txt []string, fileSave []string) {
	wordMap := make(map[string][]string)
	var save_identically_word int = 0
	var save_txt_part []string
	var randWordChoose int

	for i := 0; i < len(fileSave); i++ {
		fmt.Print(fileSave[i] + " ")
	}
	for k := len(fileSave); k < *defaultNumber; k++ {
		arrayToString := strings.Join(fileSave, " ")
		for m := 0; m < len(save_txt); m++ {
			save_txt_part = append(save_txt_part, save_txt[m])
			if len(save_txt_part) == *prefix {
				for j := 0; j < *prefix; j++ {
					if save_txt_part[j] == fileSave[j] {
						save_identically_word++
					}
				}
				if save_identically_word == *prefix {
					if m+1 < len(save_txt) {
						wordMap[arrayToString] = append(wordMap[arrayToString], save_txt[m+1])
					}
				}
			} else if len(save_txt_part) > *prefix {
				save_txt_part = save_txt_part[2:]
			}
			save_identically_word = 0
		}
		if len(wordMap[arrayToString]) == 0 {
			fmt.Println("End")
			os.Exit(1)
		}
		randWordChoose = rand.Intn(len(wordMap[arrayToString]))
		fmt.Print(wordMap[arrayToString][randWordChoose] + " ")
		fileSave = append(fileSave[1:], wordMap[arrayToString][randWordChoose])
		// fmt.Println(fileSave)
	}
	fmt.Println()
}

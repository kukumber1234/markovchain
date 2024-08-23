package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

var (
	defaultNumber = flag.Int("default number", 100, "100")
	help          = flag.Bool("help", false, "Da")
	prefix        = flag.Int("default prefix", 2, "2")
	// wordStart     = flag.String("start word", "Chapter", "First")
)

func main() {
	prefixDivide := *prefix - 1
	var allSave []string
	var fileSave []string
	var save_data []string
	var save_txt []string
	var save_word []string
	var save_prefix []string
	var randomCount int = 0
	var save_identically_word int = 0
	var save_txt_part []string
	fileName := "the_great_gatsby.txt"

	save_data = append(save_data, os.Args...)
	fileSave = fileRead(fileName)
	if len(save_data) >= 3 {
		if save_data[1] == "-w" {
			number := string_to_int(save_data[2])
			if number > 10000 {
				fmt.Fprintln(os.Stderr, "Given number can't be more 10,000")
				os.Exit(1)
			}
			*defaultNumber = number
		} else {
			fmt.Fprintln(os.Stderr, "Incorrect input")
			os.Exit(1)
		}
		if *prefix > 0 {
			for i := 0; i < *prefix; i++ {
				save_prefix = append(save_prefix, fileSave[i])
			}
		} else if *prefix == 0 {
			fmt.Fprintln(os.Stderr, "Prefix can't be equal to 0")
			os.Exit(1)
		} else if *prefix < 0 {
			fmt.Fprintln(os.Stderr, "Prefix can't be less than 0")
			os.Exit(1)
		}
		for i := 0; i < len(save_prefix); i++ {
			allSave = append(allSave, save_prefix[i])
		}
		wordMore := ""
		for {
			_, err := fmt.Fscan(os.Stdin, &wordMore)
			if err != nil {
				break
			}
			save_txt = append(save_txt, wordMore)
		}
		// save_txt_copy := save_txt
		for i := 0; i < len(save_txt); i++ {
			save_txt_part = append(save_txt_part, save_txt[i])
			if len(save_txt_part) == *prefix {
				for j := 0; j < *prefix; j++ {
					if save_prefix[j] == save_txt_part[j] {
						save_identically_word++
					}
					if save_identically_word == *prefix {
						if i+1 < len(save_txt) {
							save_word = append(save_word, save_txt[i+1])
							save_prefix = save_prefix[*prefix-prefixDivide:]
							for l := 0; l < len(save_word); l++ {
								save_prefix = append(save_prefix, save_word[l])
								allSave = append(allSave, save_word...)
								fmt.Println(save_prefix)
							}
						}
						i = 0
					}
				}
				save_identically_word = 0
				save_word = []string{}
				save_txt_part = []string{}
			}
		}
		fmt.Println(save_word)
		for i := 0; i < len(allSave); i++ {
			fmt.Print(allSave[i])
			fmt.Print(" ")
		}
		fmt.Println()
	} else {
		if len(save_data) >= 2 {
			if save_data[1] == "--help" {
				*help = true
			}
			if *help {
				helper()
			}
		}
		if *prefix > 0 {
			for i := 0; i < *prefix; i++ {
				save_txt = append(save_txt, fileSave[i])
			}
		} else if *prefix == 0 {
			fmt.Fprintln(os.Stderr, "Prefix can't be equal to 0")
			os.Exit(1)
		} else if *prefix < 0 {
			fmt.Fprintln(os.Stderr, "Prefix can't be less than 0")
			os.Exit(1)
		}
		for i := 0; i < len(save_txt); i++ {
			fmt.Printf(save_txt[i])
		}
		fmt.Printf(" ")
		wordLess := ""
		for {
			fi, _ := os.Stdin.Stat()
			if fi.Mode()&os.ModeNamedPipe == 0 {
				fmt.Println("Error: no input text")
				os.Exit(1)
			}
			_, err1 := fmt.Fscan(os.Stdin, &wordLess)
			if err1 != nil {
				break
			}
			save_txt = append(save_txt, wordLess)
		}
		inFileRandom := rand.Intn(len(save_txt))
		for i := inFileRandom; i < len(save_txt); i++ {
			fmt.Print(save_txt[i], " ")
			randomCount++
			if save_txt[i] == "." {
				randomCount--
			}
			if randomCount == *defaultNumber-*prefix {
				fmt.Println()
				os.Exit(0)
			}
		}
	}
}

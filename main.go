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
)

func main() {
	var save_data []string
	var save_txt []string
	var randomCount int = 0
	fileName := "the_great_gatsby.txt"
	fileSave := fileRead(fileName)
	save_data = append(save_data, os.Args...)
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
		wordMore := ""
		for {
			_, err := fmt.Fscan(os.Stdin, &wordMore)
			if err != nil {
				break
			}
			save_txt = append(save_txt, wordMore)
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

package main

import (
	"flag"
	"fmt"
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
	var fileSave []string
	fileName := "the_great_gatsby.txt"

	save_data = append(save_data, os.Args...)
	fileSave = fileRead(fileName)
	if len(save_data) >= 3 {
		if save_data[1] == "-w" {
			number := string_to_int(save_data[2])
			*defaultNumber = number
		} else {
			fmt.Fprintln(os.Stderr, "Incorrect input")
			os.Exit(1)
		}
		save_txt = fullFile()
		mapSave(save_txt, fileSave)
	} else {
		fi, _ := os.Stdin.Stat()
			if fi.Mode()&os.ModeNamedPipe == 0 {
				fmt.Println("Error: no input text")
				os.Exit(1)
			}
		if save_data[1] == "--help" {
			*help = true
		}
		if *help {
			helper()
		}
	}
}

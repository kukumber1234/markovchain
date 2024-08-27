package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	defaultNumber = flag.Int("w", 100, "change default number")
	help          = flag.Bool("help", false, "Check what you can do")
	prefix        = flag.Int("l", 2, "change prefix")
	prefixWord    = flag.String("p", "", "Where to start")
)

func main() {
	var save_txt []string
	var fileSave []string

	flag.Parse()

	if *prefix != len(*prefixWord) {
		fmt.Fprintln(os.Stderr, "Len of prefix word and prefix must be equal")
		os.Exit(1)
	}

	if *help {
		helper()
	}

	if *prefix != 2 {
		checkPref(prefix)
	}

	if *defaultNumber != 100 {
		checkNum(defaultNumber)
	}

	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("Error: no input text")
		os.Exit(1)
	}

	if *prefixWord != "" && *prefix != 0 {
		checkLen()
	}

	if *prefixWord != "" {
		checkWord(save_txt, fileSave)
	} else {
		save_txt = fullFile()
		fileSave = fileRead(save_txt)
		mapSave(save_txt, fileSave)
	}
}

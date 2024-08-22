package main

import (
	"fmt"
	"os"
)

func helper() {
	fmt.Println("Markov Chain text generator.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
	fmt.Println("  markovchain --help")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --help  Show this screen.")
	fmt.Println("  -w N    Number of maximum words")
	fmt.Println("  -p S    Starting prefix")
	fmt.Println("  -l N    Prefix length")
	os.Exit(0)
}

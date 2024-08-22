package main

import (
	"fmt"
	"os"
	"strconv"
)

func string_to_int(string_number string) int {
	if checkNumber(string_number) {
		if string_number[0] == '-' {
			fmt.Fprintln(os.Stderr, "Given number can't be negative.")
			os.Exit(1)
		} else if string_number[0] > 48 && string_number[0] < 57 {
			d, err := strconv.Atoi(string_number)
			if err != nil {
				os.Exit(1)
			}
			return d
		} else {
			fmt.Fprintln(os.Stderr, "Unexpected thing front of number")
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Incorrect input of number")
		os.Exit(1)
	}
	return 0
}

func checkNumber(string_number string) bool {
	for i := 1; i < len(string_number); i++ {
		if string_number[i] < 48 || string_number[i] > 57 {
			return false
		}
	}
	return true
}

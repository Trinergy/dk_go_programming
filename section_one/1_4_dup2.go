package main

import (
	"bufio"
	"fmt"
	"os"
)

// CountDuplicateInFiles (dup2) takes multiple file paths as args and counts the lines
func countDuplicatesInFiles() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}
	for _, arg := range files {
		file, err := os.Open(arg)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(file, counts)
	}
	for line, count := range counts {
		fmt.Printf("%v: %v\n", line, count)
	}
}

// countLines loads one line into a buffer and updates the counts map
func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring errors from input.Err()
}

func main() {
	countDuplicatesInFiles()
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		counts[line] = counts[line] + 1
	}
	if err := input.Err(); err != nil {
		n, err := fmt.Fprintln(os.Stderr, "reading input: ", err)
		if err != nil && n != 0 {
			fmt.Fprintln(os.Stderr, "Fprintf: %v\n", err)
			return
		}
	}
}

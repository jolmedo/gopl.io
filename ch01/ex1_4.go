package main

import (
	"bufio"
	"fmt"
	"os"
)

type WithFilenameType struct {
	amount int
	fname  string
}

func main() {
	counts := make(map[string]WithFilenameType)
	files := os.Args[1:]

	if len(files) == 0 {
		enhancedCountLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			enhancedCountLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.amount > 1 {
			fmt.Printf("%s\n", n.fname)
			fmt.Printf("%d\t%s\n", n.amount, line)
		}
	}
}

func enhancedCountLines(f *os.File, counts map[string]WithFilenameType) {
	input := bufio.NewScanner(f)
	p := WithFilenameType{0, f.Name()}
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}

		p.amount += 1
		counts[line] = p
	}
	if err := input.Err(); err != nil {
		n, err := fmt.Fprintln(os.Stderr, "reading input: ", err)
		if err != nil && n != 0 {
			fmt.Fprintln(os.Stderr, "Fprintf: %v\n", err)
			return
		}
	}
}

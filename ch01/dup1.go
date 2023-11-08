package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		counts[line] = counts[line] + 1
		//counts[input.Text()]++
	}
	if err := input.Err(); err != nil {
		n, err := fmt.Fprintln(os.Stderr, "reading input: ", err)
		if err != nil && n != 0 {
			fmt.Fprintln(os.Stderr, "Fprintf: %v\n", err)
			return
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func listingArgumentsV1(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func listingArgumentsV2(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func main() {
	var start = time.Now()
	listingArgumentsV1(os.Args)
	fmt.Printf("version1: %d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
	start = time.Now()
	listingArgumentsV2(os.Args)
	fmt.Printf("version2: %d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
}

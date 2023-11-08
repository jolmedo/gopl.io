package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", " "
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + sep + os.Args[i]
		fmt.Println(s)
		s = ""
	}

}

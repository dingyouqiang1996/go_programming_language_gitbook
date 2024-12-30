package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		sep = " "
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	elapsed := time.Since(start)
	fmt.Printf("Code took %s to execute\n", elapsed)
}

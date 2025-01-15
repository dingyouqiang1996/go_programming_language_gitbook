package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
    start := time.Now()
    var s, sep string
    for i := 0; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
	elapsed := time.Since(start)
	fmt.Printf("Code took %s to execute\n", elapsed)
}

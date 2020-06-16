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
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d %s\n", n, line)
		if n > 1 {
			fmt.Printf("big count: %d\t%s\n", n, line)
		}
	}
}

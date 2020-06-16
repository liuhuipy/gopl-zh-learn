package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 {
				fmt.Printf("file name: %s\n", filename)
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Big count: %d %s\n", n, line)
		} else {
			fmt.Printf("%d %s\n", n, line)
		}
	}
}

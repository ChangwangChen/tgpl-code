package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	//read file
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		//split and counts
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	//print
	for name, n := range counts {
		fmt.Printf("%d\t%s\n", n, name)
	}
}

package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	//for i := len(s) - 1; i >= 0; i-- {
	//	if s[i] == '/' {
	//		s = s[i+1:]
	//		break
	//	}
	//}
	//
	//for i := len(s) - 1; i >= 0; i-- {
	//	if s[i] == '.' {
	//		s = s[:i]
	//		break
	//	}
	//}

	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	s := "a/b/c.go"
	fmt.Printf("basename is: %s\n", basename(s))
}

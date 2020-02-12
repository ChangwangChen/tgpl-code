package main

import (
	"fmt"
	tc "gopl.io/tempconv"
)

func main() {
	fmt.Printf("%g\n", tc.BoilingC-tc.FreeingC)
	c := tc.FToc(100)
	fmt.Println(c.String())
	f := tc.CToF(c)
	fmt.Println(f.String())
}

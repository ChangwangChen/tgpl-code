package main

import (
	"fmt"
	"io"
	"strings"

	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Printf("fetch: reading %s %v \n", url, err)
		//	return
		//}
		//fmt.Printf("%s ", b)
		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v", err)
		}

		fmt.Printf("count %d", n)
	}
}

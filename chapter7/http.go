package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %s", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s", req.URL)
	}
}

type databases1 map[string]dollars

func (db databases1) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (db databases1) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	//db := database{"shoes": 50, "socks": 5}
	//log.Fatal(http.ListenAndServe("localhost:8000", db))

	db1 := databases1{"shoes": 50.07, "socks": 5.90}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db1.list))
	mux.Handle("/price", http.HandlerFunc(db1.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

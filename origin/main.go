package main

import (
	"net/http"
)

func main() {
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe("127.0.0.1:8087", nil); err != nil {
		panic(err)
	}
}

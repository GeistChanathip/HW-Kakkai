package main

import (
	"io"
	"net/http"
)

func main() {

	page1 := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Hello , I'm Geist!.</h1>")
	}
	page2 := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Nice to see you.</h1>")
	}

	http.HandleFunc("/", page1)
	http.HandleFunc("/about", page2)

	http.ListenAndServe(":3000", nil)

}

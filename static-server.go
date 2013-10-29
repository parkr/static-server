package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Path string
	Body []byte
}

func loadPage(path string) (*Page, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Page{Path: path, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	page, err := loadPage(r.URL.Path[1:])
	if err != nil {
		fmt.Println((err))
	}
	fmt.Fprintf(w, "%s", page.Body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

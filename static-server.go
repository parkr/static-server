package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Page struct {
	Path string
	Body []byte
}

func isDirectory(path string) bool {
	exp, _ := regexp.Compile(`\/\z`)
	return exp.Match([]byte(path))
}

func hasExtname(path string) bool {
	exp, _ := regexp.Compile(`\.\w+\z`)
	return exp.Match([]byte(path))
}

func getFilePath(thePath string) string {
	if isDirectory(thePath) {
		log.Printf(":: Adding an index.html to %s\n", thePath)
		return fmt.Sprintf("%sindex.html", thePath)
	} else {
		if !hasExtname(thePath) {
			log.Printf(":: Adding an /index.html to %s\n", thePath)
			return fmt.Sprintf("%s/index.html", thePath)
		} else {
			return thePath
		}
	}
}

func loadPage(path string) (*Page, error) {
	filePath := getFilePath(path)
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &Page{Path: filePath, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf(":: GET /%s\n", r.URL.Path[1:])
	page, err := loadPage(r.URL.Path[1:])
	if err != nil {
		log.Printf(":: Error =============== \n%s\n===============\n", err)
		fmt.Fprintf(w, "There was an error fetching that. Try again?")
	} else {
		fmt.Fprintf(w, "%s", page.Body)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

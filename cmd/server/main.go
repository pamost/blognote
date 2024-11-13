package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pamost/blognote"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	mux := http.NewServeMux()

	postReader := blognote.FileReader{Dir: "posts"}

	postTemplate := template.Must(template.ParseFiles("post.html"))
	mux.HandleFunc("GET /post/{slug}", blognote.PostHandler(postReader, postTemplate))

	indexTemplate := template.Must(template.ParseFiles("index.html"))
	mux.HandleFunc("GET /", blognote.IndexHandler(postReader, indexTemplate))

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

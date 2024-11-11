package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {  
	mux := http.NewServeMux()

	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReaded{}))

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type SlugReader interface {
	Read(slug string) (string, error)
}
type FileReaded struct{}

func (fr FileReaded) Read(slug string) (string, error) {
	f, err := os.Open(slug + ".md")
	if err != nil {
		return "", nil
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}

func PostHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postmarkDown, err := sl.Read(slug)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, postmarkDown)
	}
}


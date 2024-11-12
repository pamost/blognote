package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func main() {
	mux := http.NewServeMux()

	postTemplate := template.Must(template.ParseFiles("post.html"))
	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReaded{}, postTemplate))

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

type PostData struct {
	Content string
	Author  string
	Title   string
}

func PostHandler(sl SlugReader, tpl *template.Template) http.HandlerFunc {
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)

	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postmarkDown, err := sl.Read(slug)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		var buf bytes.Buffer
		err = md.Convert([]byte(postmarkDown), &buf)
		if err != nil {
			panic(err)
		}

		err = tpl.Execute(w, PostData{
			Content: buf.String(),
			Author:  "Pavel",
			Title:   "BlogNote",
		})
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	}
}

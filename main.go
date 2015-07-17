package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type Data struct {
	Title string
	Body  template.HTML
}

func (d Data) Files(types ...string) []string {
	files := []string{}
	for _, t := range types {
		fs, _ := filepath.Glob("files/*." + t)
		files = append(files, fs...)
	}
	return files
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	path := "index"
	if r.URL.Path != "/" {
		path = strings.Trim(r.URL.Path, "/")
	}
	f, _ := ioutil.ReadFile(path + ".md")
	lines := strings.Split(string(f), "\n")
	templateFile := strings.TrimSpace(lines[0])
	markdown := strings.Join(lines[1:], "\n")
	body := blackfriday.MarkdownCommon([]byte(markdown))

	d := Data{
		Title: "Welcome!",
		Body:  template.HTML(body),
	}

	t := template.New(templateFile)
	t, _ = template.ParseFiles(templateFile)
	t.Execute(w, d)
}

func getFiles(path string) []string {
	files, _ := filepath.Glob(path + "*")
	return files
}

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/Users/colsen/Projects/ktc/files"))))
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":8000", nil)
}

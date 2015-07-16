package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

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

	d := struct {
		Title string
		Items []string
		Body  template.HTML
	}{
		Title: "Welcome!",
		Items: []string{"item1", "item2", "item3"},
		Body:  template.HTML(body),
	}

	t := template.New(templateFile)
	t, _ = template.ParseFiles(templateFile)
	t.Execute(w, d)
}

func main() {
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":8000", nil)
}

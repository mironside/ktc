package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
)

func handlerequest(w http.ResponseWriter, r *http.Request) {
	f, _ := ioutil.ReadFile("index.md")
	md := blackfriday.MarkdownCommon(f)

	d := struct {
		Title string
		Items []string
		Body  template.HTML
	}{
		Title: "Welcome!",
		Items: []string{"item1", "item2", "item3"},
		Body:  template.HTML(md),
	}

	t := template.New("index.html")
	t, _ = template.ParseFiles("index.html")
	t.Execute(w, d)
}

func main() {
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":8000", nil)
}

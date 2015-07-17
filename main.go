package main

import (
	//"github.com/russross/blackfriday"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	Dir   string
	Title string
	Body  template.HTML
}

func (d Data) Files(types ...string) []string {
	files := []string{}
	for _, t := range types {
		path := "files/*." + t
		if len(d.Dir) > 0 {
			path = d.Dir + "/" + path
		}
		fmt.Println(path)
		fs, _ := filepath.Glob(path)
		for _, f := range fs {
			f = strings.Trim(strings.Replace(f, "\\", "/", -1), d.Dir+"/")
			fmt.Println(f)
			files = append(files, f)
		}
	}
	return files
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	path := "index"
	if r.URL.Path != "/" {
		path = strings.Trim(r.URL.Path, "/")
	}
	src, err := os.Stat(path)
	if err == nil && src.IsDir() {
		path += "/index.md"
	} else {
		path += ".md"
	}

	fmt.Println(path)
	f, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(f), "\n")
	templateFile := strings.TrimSpace(lines[0])
	markdown := strings.Join(lines[1:], "\n")
	body := markdown

	d := Data{
		Dir:   filepath.Dir(path),
		Title: "Welcome!",
		Body:  template.HTML(body),
	}

	t := template.New(templateFile)
	t, _ = template.ParseFiles(templateFile)
	bt := t.New("body")
	bt, _ = bt.Parse(body)

	t.Execute(w, d)
}

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("C:/temp/ktc/files"))))
	http.Handle("/john/files/", http.StripPrefix("/john/files/", http.FileServer(http.Dir("C:/temp/ktc/john/files"))))
	http.HandleFunc("/", handlerequest)
	http.ListenAndServe(":8000", nil)
}

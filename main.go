package main

import (
	"fmt"
	"html/template"
	// "os"
	"net/http"
)

var tpl *template.Template

//init initalizes the html templates
func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println(err)
	}
}

//index
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hello.html", nil)
}

//processor
func processor(w http.ResponseWriter, r *http.Request) {
	// GET sends this through the url and POST doesn't
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")

	d := struct {
		First string
	}{
		First: fname,
	}

	fmt.Printf("%v\n", fname)
	tpl.ExecuteTemplate(w, "processor.html", d)
}

//Webhello
func Webhello(w http.ResponseWriter, r *http.Request) {
	name := Name{"Gavin", "Jaeger-Freeborn"}
	template, _ := template.ParseFiles("hello.html")
	template.Execute(w, name)
}

type Name struct {
	FName, LName string
}

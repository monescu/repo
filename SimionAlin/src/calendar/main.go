package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var date = r.FormValue("date")
	t, _ := time.Parse("2006-01-02", date)
	body := []byte(date + " " + t.Weekday().String())
	p := &Page{Title:"Date",
		Body: body}
	renderTemplate(w, "view", p)
}

var templates = template.Must(template.ParseFiles("view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(":8880", nil))
}
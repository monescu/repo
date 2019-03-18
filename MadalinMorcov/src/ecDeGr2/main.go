package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func equationHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "equation", p)
}

func solveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))
	c, _ := strconv.Atoi(r.FormValue("c"))
	delta := float64(b*b - 4*a*c)

	if delta < 0 {
		body = "Negative delta"
	} else if delta == 0 {
		result := (-b)/(2*a)
		body = fmt.Sprintf("%.2f", result)
	} else {
		result1 := (float64(-b) + math.Sqrt(delta)) / (2*float64(a))
		result2 := (float64(-b) - math.Sqrt(delta)) / (2*float64(a))
		buffer := bytes.Buffer{}
		buffer.WriteString("X1: ")
		buffer.WriteString(fmt.Sprintf("%.2f", result1))
		buffer.WriteString("\nX2: ")
		buffer.WriteString(fmt.Sprintf("%.2f", result2))
		body = buffer.String()
	}
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


var templates = template.Must(template.ParseFiles("view.html", "equation.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view|solve|equation)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/solve/", makeHandler(solveHandler))
	http.HandleFunc("/equation/", makeHandler(equationHandler))

	log.Fatal(http.ListenAndServe(":8880", nil))
}
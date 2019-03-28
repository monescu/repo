package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"timetable"
)

var courses []timetable.Course

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/timetable", timetableHandler)

	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8990", nil); err != nil {
		log.Fatal(err)
	}
}

func timetableHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/timetable" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		courses = timetable.GetCourses()
		searchTemplate := timetable.GetSearchTemplate(courses)
		timetableTemplate := template.Must(template.ParseFiles("static/timetableForm.html"))

		timetableTemplate.Execute(w, searchTemplate)

		//http.ServeFile(w, r, "static/timetable.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		teacher := r.FormValue("teacher")
		day := r.FormValue("day")
		discipline := r.FormValue("discipline")
		year := r.FormValue("year")
		specialisation := r.FormValue("specialisation")
		courseType := r.FormValue("courseType")
		group := r.FormValue("group")
		semiGroup := r.FormValue("semiGroup")
		hours := r.FormValue("hours")

		courses = timetable.GetCourses()

		searchTemplate := timetable.GetSearchTemplate(courses)
		searchTemplate.IsPost = true
		searchTemplate.MatchedCourses = timetable.Search(courses, teacher, day, discipline, year, specialisation, courseType, group, semiGroup, hours)

		timetableTemplate := template.Must(template.ParseFiles("static/timetableForm.html"))

		timetableTemplate.Execute(w, searchTemplate)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

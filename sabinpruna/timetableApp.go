package main

import (
	"fmt"
	"net/http"

	"github.com/tealeg/xlsx"
)

func main() {

	unmerged, _ := xlsx.FileToSliceUnmerged("../Orar_sem_II_2018-2019_V4.xlsx")

	fmt.Println(unmerged)

	// http.Handle("/", http.FileServer(http.Dir("./static")))

	// http.HandleFunc("/timetable", timetableHandler)

	// fmt.Printf("Starting server for testing HTTP POST...\n")
	// if err := http.ListenAndServe(":8990", nil); err != nil {
	// 	log.Fatal(err)
	// }
}

func timetableHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/timetable" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/timetable.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

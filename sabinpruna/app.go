package main

import (
	"encoding/gob"
	"fmt"
	"homework"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := [9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9}}
	fmt.Println(homework.PrintSudokuBoard(numbers))

	message := "Servusi"
	fmt.Println(homework.MatrixCipher(message))

	message = "eSvrsui"
	fmt.Println(homework.MatrixCipher(message))

	http.HandleFunc("/quadratic", quadraticHandler)
	http.HandleFunc("/sudoku", sudokuSolverHandler)
	http.HandleFunc("/circle", circleHandler)

	//This can be simplified if i find out how to check who sent the request
	//probably use templates and have a field for the action executed
	http.HandleFunc("/caesar", caesarHandler)
	http.HandleFunc("/caesarDecrypt", caesarDecryptHandler)
	http.HandleFunc("/caesarUpload", caesarUploadHandler)

	//timetable
	///
	//http.HandleFunc("/timetable", timetableHandler)

	//1 http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//2 mux := http.NewServeMux()
	// mux.Handle("/", http.FileServer(http.Dir("./static")))
	//3 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	//fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	// })

	//Somehow this fucking magic worked and not the others
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8990", nil); err != nil {
		log.Fatal(err)
	}

}

func quadraticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/quadratic" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/quadraticForm.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		a, _ := strconv.Atoi(r.FormValue("a"))
		b, _ := strconv.Atoi(r.FormValue("b"))
		c, _ := strconv.Atoi(r.FormValue("c"))

		complexA := complex(float64(a), 0)
		complexB := complex(float64(b), 0)
		complexC := complex(float64(c), 0)
		positiveSolution, negativeSolution := homework.SolveQuadratic(complexA, complexB, complexC)

		fmt.Fprintf(w, "Positive Solution = %v\n", positiveSolution)
		fmt.Fprintf(w, "Negative Solution = %v\n", negativeSolution)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func sudokuSolverHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sudoku" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/sudokuForm.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		var sudokuNumbers [9][9]int

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				sudokuNumbers[i][j], _ = strconv.Atoi(r.FormValue(fmt.Sprintf("cell-%d%d", i, j)))

			}
		}

		if ok := homework.BacktrackSudokuSolution(&sudokuNumbers); ok {
			fmt.Fprintf(w, homework.PrintSudokuBoard(sudokuNumbers))
		} else {
			fmt.Fprintf(w, "Invalid configuration!")
		}
	}
}

func circleHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/circle" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/circleForm.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		counter := 0
		var points []homework.Point
		for index := range r.Form {
			counter++
			//can't think right now, just used it so no compiler suggestions
			fmt.Println(index)
		}

		for index := 1; index <= counter; index++ {
			var p homework.Point
			temp := strings.Fields(r.FormValue(fmt.Sprintf("point-%d", index))) //values[0]

			p.X, _ = strconv.Atoi(temp[0])
			p.Y, _ = strconv.Atoi(temp[1])
			points = append(points, p)
		}

		radius := homework.FindSmallestCircle(points)

		/**
		 * Alternative with matrix :
		 * radius, center := homework.FindSMallestCircleMatrix(points)
		 *
		 * output := "The smallest circle with centre(%d,%d) that contains all the points has the radius: %d"
		 * fmt.Fprintf(w, fmt.Sprintf(output, center.X, center.Y, radius))
		 **/

		output := "The smallest circle with centre(%d,%d) that contains all the points has the radius: %d"
		fmt.Fprintf(w, fmt.Sprintf(output, points[0].X, points[0].Y, radius))
	}
}

func caesarHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/caesar" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		file := homework.GenerateRandomNumberFile()

		//http.ServeFile(w, r, "static/caesar.html")

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", file.Name()))
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

		//http.ServeContent(w, r, file.Name(), file)
		http.ServeFile(w, r, file.Name())
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		r.ParseMultipartForm(32 << 20)

		// parse and validate file and post parameters
		file, _, err := r.FormFile("key")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE"))
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE"))
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "text/plain; charset=utf-8":
			break
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE_TYPE"))
			return
		}

		numbersString := string(fileBytes)
		numberStringSlice := strings.Fields(numbersString)
		numbers := []int{}

		for _, numberString := range numberStringSlice {
			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}

		message := r.FormValue("message")

		message = homework.Encrypt(numbers, message)

		binaryFile, binaryErr := os.Create("test.bin")
		defer binaryFile.Close()

		if binaryErr != nil {
			log.Fatal(binaryErr)
		}

		// var binaryBuffer bytes.Buffer
		// binary.Write(&binaryBuffer, binary.BigEndian, message)
		// binaryFile.Write(binaryBuffer.Bytes())
		//binaryBuffer := bufio.NewWriter(binaryFile)
		encoder := gob.NewEncoder(binaryFile)

		if err := encoder.Encode(message); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, message)
	}

}

func caesarUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/caesarUpload" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "static/caesar.html")

}

func caesarDecryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/caesarDecrypt" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/caesarDecrypt.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		r.ParseMultipartForm(32 << 20)

		// parse and validate file and post parameters
		file, _, err := r.FormFile("key")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE"))
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE"))
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "text/plain; charset=utf-8":
			break
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE_TYPE"))
			return
		}

		numbersString := string(fileBytes)
		numberStringSlice := strings.Fields(numbersString)
		numbers := []int{}

		for _, numberString := range numberStringSlice {
			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}

		message := r.FormValue("message")

		message = homework.Decrypt(numbers, message)

		binaryFile, err := os.Open("test.bin")
		defer binaryFile.Close()

		if err != nil {
			log.Fatal(err)
		}

		var binaryMessage string
		decoder := gob.NewDecoder(binaryFile)
		decoder.Decode(&binaryMessage)
		fmt.Println(homework.Decrypt(numbers, binaryMessage))

		fmt.Fprintf(w, message)
	}

}

func timetableHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/timetable" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/timetableForm.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// parse and validate file and post parameters
		file, _, err := r.FormFile("key")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE"))
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE"))
			return
		}

		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		switch filetype {
		case "text/plain; charset=utf-8":
			break
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("INVALID_FILE_TYPE"))
			return
		}

		numbersString := string(fileBytes)
		numberStringSlice := strings.Fields(numbersString)
		numbers := []int{}

		for _, numberString := range numberStringSlice {
			number, _ := strconv.Atoi(numberString)
			numbers = append(numbers, number)
		}

		message := r.FormValue("message")

		message = homework.Decrypt(numbers, message)

		fmt.Fprintf(w, message)
	}

}

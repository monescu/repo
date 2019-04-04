package concurrency

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//owing to github/habib-rangoonwala

// collector
type mapCollector chan chan interface{}

// mapper
type mapperFunction func(interface{}, chan interface{})

// Reducer
type reducerFunction func(chan interface{}, chan interface{})

const (
	//MaxThreads is max for my processor
	MaxThreads = 8
)

//ParallelWordCount  returns word count of directory using goroutines
func ParallelWordCount() {
	starttime := time.Now()
	fmt.Println("Processing . . .")

	input := getFilesConcurrent("D:\\Facultate\\go_course\\repo\\sabinpruna")

	// this will start the map reduce work
	results := mapReduce(mapper, reducer, input)

	filehandle, err := os.Create("D:\\Facultate\\go_course\\repo\\sabinpruna" + "/processed.csv")
	defer filehandle.Close()
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	fmt.Println("Writing to file:", filehandle.Name())

	writer := bufio.NewWriter(filehandle)

	for word, frequency := range results.(map[string]int) {
		fmt.Fprintln(writer, word+","+strconv.Itoa(frequency))
	}
	defer writer.Flush()
	defer filehandle.Close()

	elapsedtime := time.Since(starttime)
	fmt.Println("Complete")
	fmt.Println("Time taken:", elapsedtime)
}

//-------------------------PRIVATES-------------------------------------------
func getFilesConcurrent(dirname string) chan interface{} {
	output := make(chan interface{})
	go func() {
		filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				output <- path
			}
			return nil
		})
		close(output)
	}()
	return output
}

func parseFileConcurrent(filename string) chan string {
	output := make(chan string)
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")

	go func() {
		fmt.Println("Reading file : " + filename + "\n")
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			return
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			word := strings.TrimSpace(reg.ReplaceAllString(scanner.Text(), ""))
			if len(word) > 0 {
				output <- word
			}
		}

		close(output)
		fmt.Println("Completed file : " + filename + "\n")

	}()
	return output
}

//----------------------------------------------------------------------------

//-------------------------MAP REDUCE-----------------------------------------

func reducerDispatcher(collector mapCollector, reducerInput chan interface{}) {

	for output := range collector {
		reducerInput <- <-output
	}
	close(reducerInput)
}

func mapperDispatcher(mapper mapperFunction, input chan interface{}, collector mapCollector) {

	for item := range input {
		taskOutput := make(chan interface{})
		go mapper(item, taskOutput)
		collector <- taskOutput

	}
	close(collector)
}

func mapper(filename interface{}, output chan interface{}) {
	results := map[string]int{}

	for word := range parseFileConcurrent(filename.(string)) {
		results[strings.ToLower(word)]++
	}

	output <- results
}

func reducer(input chan interface{}, output chan interface{}) {

	results := map[string]int{}
	for matches := range input {
		for word, frequency := range matches.(map[string]int) {
			results[strings.ToLower(word)] += frequency
		}
	}
	output <- results
}

func mapReduce(mapper mapperFunction, reducer reducerFunction, input chan interface{}) interface{} {

	reducerInput := make(chan interface{})
	reducerOutput := make(chan interface{})
	MapCollector := make(mapCollector, MaxThreads)

	go mapperDispatcher(mapper, input, MapCollector)
	go reducerDispatcher(MapCollector, reducerInput)
	go reducer(reducerInput, reducerOutput)

	return <-reducerOutput
}

//------------------------------------------------------------------------

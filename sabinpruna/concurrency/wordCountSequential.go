package concurrency

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//SequentialWordCount parses files in directory in a sequential fashion
func SequentialWordCount() {

	starttime := time.Now()
	fmt.Println("Processing . . .")

	words := map[string]int{}

	files := getFilesFromDirectory("D:\\Facultate\\go_course\\repo\\sabinpruna")

	for _, file := range files {
		parseFile(file, words)
	}

	printWords(words, "D:\\Facultate\\go_course\\repo\\sabinpruna")
	elapsedtime := time.Since(starttime)
	fmt.Println("Complete")
	fmt.Println("Time taken:", elapsedtime)

}

//---------------------PRIVATES-------------------------------------------------

func getFilesFromDirectory(searchDir string) []string {
	files := []string{}

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func parseFile(filename string, words map[string]int) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	fmt.Println(". . . " + filename)

	scanner := bufio.NewScanner(file)

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.TrimSpace(reg.ReplaceAllString(scanner.Text(), ""))
		if len(word) > 0 {
			words[strings.ToLower(word)]++
		}
	}
}

func printWords(words map[string]int, filename string) {

	filehandle, err := os.Create(filepath.Dir(filename) + "/processed.csv")
	if err != nil {
		log.Println("Error writing to file: ", err)
		return
	}
	fmt.Println("Writing to file:", filehandle.Name())

	writer := bufio.NewWriter(filehandle)

	for word, frequency := range words {
		fmt.Fprintln(writer, word+","+strconv.Itoa(frequency))
	}

	defer writer.Flush()
	defer filehandle.Close()
}
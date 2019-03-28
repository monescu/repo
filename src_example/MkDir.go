package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// create a TestDir directory on current working directory
	//os.Mkdir("." + string(filepath.Separator) + "TestDir",0777)

	fileNames, errNames := os.Open("Studenti.txt")

	fileEmails, errEmails := os.Open("Emails.txt")

	if errNames != nil {
		log.Fatal(errNames)
	}
	defer fileNames.Close()

	if errNames != nil {
		log.Fatal(errEmails)
	}
	defer fileEmails.Close()


	scannerNames := bufio.NewScanner(fileNames)
	scannerEmails := bufio.NewScanner(fileEmails)


	for scannerNames.Scan() {
		scannerEmails.Scan()
		//fmt.Println(scannerNames.Text(), scannerEmails.Text())
		os.Mkdir("." + string(filepath.Separator) + scannerNames.Text(),0777)
	}

}
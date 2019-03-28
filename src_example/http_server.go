package main

import (
	//"bufio"
	"fmt"
	"log"
	"net/http"
	//"math"
	//"os"
)


//var users int

const msg = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
Request no %v <br>
Method %v
</body>
</html>
`



func main() {
	//stdin read
	/*
	  reader := bufio.NewReader(os.Stdin)
	  fmt.Print("Enter text: ")
	  text, _ := reader.ReadString('\n')
	  fmt.Println(text)

	  fmt.Println("Enter text: ")
	  text2 := ""
	  fmt.Scanln(&text2)
	  fmt.Println(text2)

	  var a string

	  //?cnt
	  cnt, err:=fmt.Scanln("%v", &a)
	  fmt.Println(cnt, err)
	*/

	/////////////////////////////////////////////////////////

	http.HandleFunc("/test", testHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func testHandleFunc(w http.ResponseWriter, r *http.Request){
	fmt.Println("Requesting...")
	//users++
	s:=fmt.Sprintf(msg,r.Method)
	fmt.Fprint(w, s)
}
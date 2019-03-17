package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a file and use bufio.NewWriter.
	f, _ := os.Create("file.txt")   //C:\\programs\\file.txt")
	w := bufio.NewWriter(f)

	// Use Fprint to write things to the file.
	// ... No trailing newline is inserted.
	fmt.Fprint(w, "Hello")
	fmt.Fprint(w, 123)
	fmt.Fprint(w, "...")

	// Use Fprintf to write formatted data to the file.
	value1 := "cat"
	value2 := 900
	fmt.Fprintf(w, "%v %d...", value1, value2)

	fmt.Fprintln(w, "DONE...")

	// Done.
	w.Flush()
}
package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	input := "   first   second       third    "
	scanner := bufio.NewScanner(strings.NewReader(input))

	//
	//reader:=bufio.NewReader(os.Stdin)
	//input, _:=reader.ReadString('\n')
    //

	//scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
package main

import "fmt"
import "bufio"
import "os"

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter:")
	input, err = inputReader.ReadString('q')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

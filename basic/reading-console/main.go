package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Works for 1st and 2nd Examples
	// reader := bufio.NewReader(os.Stdin)

	scanner()

	// // First Example
	// fmt.Println("Simple Shell")
	// fmt.Println("--------------------")

	// for {
	// 	fmt.Println("->")
	//// Where /n is the delimiter
	// 	text, _ := reader.ReadString('\n')
	// 	// convert CRLF to CL
	// 	text = strings.Replace(text, "\n", "", -1)

	// 	if strings.Compare("hi", text) == 0 {
	// 		fmt.Println("hello, Yourself")
	// 	}
	// }

	// // Second example

	// fmt.Println("Press a or A key\n--->")
	// char, _, err := reader.ReadRune()

	// if err != nil {
	// 	panic(err.Error())
	// }
	// // print out the unicode value i.e. A -> 65, a -> 97
	// fmt.Println(char)

	// switch char {
	// case 'A':
	// 	fmt.Println("A key pressed")
	// 	break
	// case 'a':
	// 	fmt.Println("a key pressed")
	// 	break
	// }

	// Third Example
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("This is what you wrote in the input: ", scanner.Text())
	}
	// The above code will infinitely ask scan for input and echo back whatever is entered.
}

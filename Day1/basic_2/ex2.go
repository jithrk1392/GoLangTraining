package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter a word \n")
	fmt.Scan(&input)
	for _, char := range input {
		fmt.Print(string(char), string(char))
	}
	fmt.Println("")
}

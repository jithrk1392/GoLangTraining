package main

import "fmt"

func main() {

	var factorialNum, fact int
	fact = 1

	fmt.Print("Enter any Number to find the Factorial = ")
	fmt.Scanln(&factorialNum)

	for i := 1; i <= factorialNum; i++ {
		fact = fact * i
	}
	fmt.Println("The Factorial of ", factorialNum, " = ", fact)
}

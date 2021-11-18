package main

import "fmt"

func main() {
	var n int
	t1 := 0
	t2 := 1
	next := 0

	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&n)
	fmt.Print("Fibonacci Series :")
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Print(" ", t1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", t2)
			continue
		}
		next = t1 + t2
		t1 = t2
		t2 = next
		fmt.Print(" ", next)
	}
	fmt.Print("\n")
}

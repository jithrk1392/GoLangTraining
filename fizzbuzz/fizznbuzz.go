package main

import (
	"fmt"

	"github.com/AgarwalConsulting/Go-Training/Day1extrafries/fizzbuzz/fbpak"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fbpak.FizzBuzz(i + 1))
	}
}

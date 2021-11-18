package main

import (
	"fmt"

	"jitr.package/fbpak/fbpak"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fbpak.FizzBuzz(i + 1))
	}
}

package fbpak

import (
	"fmt"
	"strconv"
)

func FizzBuzz(number int) string {
	for i := 1; i <= 100; i++ {
		isDivBy3 := i%3 == 0
		isDivBy5 := i%5 == 0

		if isDivBy3 && isDivBy5 {
			fmt.Printf("fizzbuzz")
		} else if isDivBy3 {
			fmt.Printf("fizz")
		} else if isDivBy5 {
			fmt.Printf("buzz")
		} else {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}

	return strconv.Itoa(number)

}

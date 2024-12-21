package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number < 1 || number > 10000 {
		return
	}

	tens := (number / 10) % 10

	fmt.Println(tens)
}

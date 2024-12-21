package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)

	if d <= 0 || d >= 360 {
		return
	}

	hours := d / 30
	minutes := (d % 30) * 2

	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}

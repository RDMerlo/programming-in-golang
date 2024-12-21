package main

import "fmt"

func main() {
	var R float64
	fmt.Scan(&R)
	if R <= 0 {
		fmt.Printf("число %2.2f не подходит", R)
	} else if R > 10000 {
		fmt.Printf("%e", R)
	} else {
		fmt.Printf("%.4f", R*R)
	}
}

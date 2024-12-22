package main

import (
	"fmt"
	"math"
)

var k float64 = 1296
var p float64 = 6
var v float64 = 6

func main() {

	fmt.Println(T())
}

func M() float64 {
	m := p * v
	return m
}

func W() float64 {
	m := M()
	w := math.Sqrt(k / m)
	return w

}

func T() float64 {
	w := W()
	t := 6 / w
	return t
}

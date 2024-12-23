/*
Входные данные
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
Выходные данные
Если треугольник существует, выведите строку "Существует", иначе выведите строку "Не существует". Строку выводите без кавычек.
*/
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b > c && b+c > a && a+c > b {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}

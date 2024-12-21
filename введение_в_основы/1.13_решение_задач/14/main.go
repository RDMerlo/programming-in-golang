/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.
Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.
*/
package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fib := []int{0, 1}
	for i := 2; ; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
		if a == fib[i] {
			fmt.Println(i)
			break
		}
		if a < fib[i] {
			fmt.Println(-1)
			break
		}
	}

}

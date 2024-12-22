// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(num uint) uint {
		var s string
		a := strconv.FormatUint(uint64(num), 10)
		for _, v := range a {
			digit, _ := strconv.Atoi(string(v))
			if digit%2 == 0 && digit != 0 {
				s += string(v)
			}
		}
		result, _ := strconv.ParseUint(s, 10, 64)
		if result == 0 {
			return uint(100)
		}
		return uint(result)
	}

	fmt.Println(fn(727178))
	fmt.Println(fn(0))
}

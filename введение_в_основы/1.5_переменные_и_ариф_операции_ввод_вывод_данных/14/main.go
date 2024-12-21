package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	// Проверка на диапазон
	if number < 1 || number > 10000 {
		return
	}

	// Получение последней цифры
	lastDigit := number % 10

	// Вывод результата
	fmt.Println(lastDigit)
}

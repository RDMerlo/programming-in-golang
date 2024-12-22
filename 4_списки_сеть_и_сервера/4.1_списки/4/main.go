package main

import (
	"container/list"
	"fmt"
)

// Push добавляет элемент в конец очереди
func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

// Pop удаляет элемент из начала очереди и возвращает его
func Pop(queue *list.List) interface{} {
	if queue.Len() == 0 {
		return nil // Очередь пуста
	}
	elem := queue.Front() // Получаем первый элемент
	queue.Remove(elem)    // Удаляем его из очереди
	return elem.Value     // Возвращаем значение
}

// printQueue печатает очередь в одну строку без пробелов
func printQueue(queue *list.List) {
	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		fmt.Print(elem.Value)
	}
	fmt.Println()
}

func main() {

	queue := list.New()

	Push(1, queue)

	Push(2, queue)

	Push(3, queue)

	printQueue(queue) // 123

	Pop(queue)

	printQueue(queue) // в ту же строку: 23

}

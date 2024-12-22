package main

import (
	"container/list"
	"fmt"
)

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	reversedList := list.New() // Создаем новый список для результата
	for elem := l.Back(); elem != nil; elem = elem.Prev() {
		reversedList.PushBack(elem.Value)
	}
	return reversedList
}

func printList(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
	fmt.Println()
}

func main() {

	task := list.New()

	for i := 0; i < 10; i++ {
		task.PushBack(i)
	}
	printList(task)
	// 0123456789

	printList(ReverseList(task))
	// 9876543210

}

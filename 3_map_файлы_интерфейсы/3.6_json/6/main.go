/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Rating     []int
}

type groupStudents struct {
	Id       int
	Number   string
	Year     int
	Students []student
}

type averageJSON struct {
	Average float64
}

func main() {
	var gs groupStudents

	//считываем из станд.ввода
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	//декодируем данные в перем gs
	if err := json.Unmarshal(data, &gs); err != nil {
		fmt.Println(err)
	}

	//считаем кол-во оценок
	var countRating = 0
	for _, arrStudents := range gs.Students {
		countRating = countRating + len(arrStudents.Rating)
	}

	//среднее количество оценок, полученное студентами группы
	average := float64(countRating) / float64(len(gs.Students))

	//создаем выходную структуру JSON
	out := averageJSON{
		Average: average,
	}
	//кодируем в JSON с форматированием вывода
	dataJSON, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	//выводим в стандартный вывод
	os.Stdout.Write(dataJSON)

}

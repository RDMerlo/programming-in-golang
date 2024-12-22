package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

func main() {
	http.HandleFunc("/api/user", handleUserRequest)
	//log.Println("Server starting on port 9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		//log.Fatalf("Failed to start server: %v", err)
	}
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	// Получаем значение параметра name из запроса
	name := r.URL.Query().Get("name")

	// Формируем ответ
	response := fmt.Sprintf("Hello,%s!", name)

	// Отправляем ответ клиенту
	w.Write([]byte(response))
}

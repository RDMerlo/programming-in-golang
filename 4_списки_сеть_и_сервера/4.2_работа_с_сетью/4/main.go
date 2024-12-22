package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024) // Буфер для чтения данных

	for i := 0; i < 3; i++ {
		// Читаем данные от сервера
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Ошибка чтения данных: %v", err)
		}

		// Преобразуем данные в строку и выводим в верхнем регистре
		message := strings.ToUpper(string(buffer[:n]))
		fmt.Println(message)
	}
}

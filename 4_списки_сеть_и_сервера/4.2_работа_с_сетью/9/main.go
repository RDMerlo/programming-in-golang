package main

// данные пакеты нужны для системы проверки
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Устанавливаем таймаут для HTTP клиента
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Выполняем HTTP GET запрос
	resp, err := client.Get("http://127.0.0.1:5555/get")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка запроса: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения тела ответа: %v\n", err)
		os.Exit(1)
	}

	// Печатаем тело ответа
	fmt.Println(string(body))
}

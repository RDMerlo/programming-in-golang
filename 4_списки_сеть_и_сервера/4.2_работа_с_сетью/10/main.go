package main

// данные пакеты нужны для системы проверки
import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// Считываем имя и возраст с консоли
	var name string
	var age string
	fmt.Scanln(&name)
	fmt.Scanln(&age)

	// Формируем URL с query параметрами
	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Устанавливаем таймаут для HTTP клиента
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Выполняем HTTP GET запрос
	resp, err := client.Get(fullURL)
	if err != nil {
		fmt.Printf("Ошибка запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения тела ответа: %v\n", err)
		return
	}

	// Печатаем тело ответа
	fmt.Println(string(body))
}

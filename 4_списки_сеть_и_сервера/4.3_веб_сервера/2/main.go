package main

// некоторые импорты нужны для проверки
import (
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello, web!"
		w.Write([]byte(message))
	})

	//fmt.Println("Server is listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		//fmt.Printf("Server failed to start: %v\n", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем статус и тип содержимого
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	// Отправляем ответ
	fmt.Fprintf(w, "Hello! You accessed %s\n", r.URL.Path)
}

func main() {
	// Регистрируем обработчик для всех путей
	http.HandleFunc("/", handler)

	// Запускаем сервер на порту 8080
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

// обработчик для главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Добро пожаловать на главную страницу!")
}

// обработчик для страницы /about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это страница о нас.")
}

func main() {
	http.HandleFunc("/", homeHandler)       // маршрут для главной страницы
	http.HandleFunc("/about", aboutHandler) // маршрут для страницы о нас

	// запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

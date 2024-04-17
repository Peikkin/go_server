package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	RunServer(":8080", nil)
}

func RunServer(addr string, handler http.Handler) {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка при запуске сервера")
	}
	log.Info().Msg("Запуск сервера на порту 8080")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/hello" {
		http.Error(w, "путь не найден", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Привет!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ошибка ParseForm: %v\n", err)
		return
	}

	fmt.Fprintf(w, "успешный POST запрос\n")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "Имя: %s\n", name)
	fmt.Fprintf(w, "Адрес: %s\n", adress)
}

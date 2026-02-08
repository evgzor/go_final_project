package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/evgzor/go_final_project/pkg/api"
)

const webDir = "./web"

const defaultPort = "7540"

const TODO_PORT = "TODO_PORT"

// Run запускает HTTP-сервер приложения.
//
// Сервер:
//   - читает порт из переменной окружения TODO_PORT
//   - использует defaultPort, если переменная не задана
//   - инициализирует API
//   - раздаёт статические файлы из webDir
func Run() {
	port := os.Getenv(TODO_PORT)

	if port == "" {
		port = defaultPort
	}

	port = ":" + port

	api.Init()

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

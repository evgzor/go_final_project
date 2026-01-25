package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/evgzor/go_final_project/pkg/api"
)

const webDir = "./web"

const TODO_PORT = "TODO_PORT"

func Run() {
	port := os.Getenv(TODO_PORT)

	if port == "" {
		port = "7540"
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

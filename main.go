package main

import (
	"net/http"
	"os"
)

const webDir = "./web"

var port = os.Getenv("TODO_PORT")

// func mainHandle(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "answer")
// }

func main() {

	if port == "" {
		port = "7540"
	}

	port = ":" + port

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

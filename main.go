package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/evgzor/go_final_project/pkg/db"
)

const webDir = "./web"

var port = os.Getenv("TODO_PORT")

// func mainHandle(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "answer")
// }

func main() {

	err := db.Init("scheduler.db")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if port == "" {
		port = "7540"
	}

	port = ":" + port

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

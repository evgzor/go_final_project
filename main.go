package main

import (
	"fmt"

	"github.com/evgzor/go_final_project/pkg/db"
	"github.com/evgzor/go_final_project/pkg/server"
	"github.com/joho/godotenv"
)

const DB_FILE_NAME = "scheduler.db"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Init(DB_FILE_NAME)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	server.Run()

	defer db.CloseDb()
}

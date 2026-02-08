package main

import (
	"fmt"

	"github.com/evgzor/go_final_project/pkg/db"
	"github.com/evgzor/go_final_project/pkg/server"
	"github.com/joho/godotenv"
)

const dbFileName = "scheduler.db"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Init(dbFileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.CloseDb()
	server.Run()
}

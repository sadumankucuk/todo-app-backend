package main

import (
	"fmt"
	"log"
	"os"
	"todo/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	svr := server.NewServer()
	err := svr.StartServer(port)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}

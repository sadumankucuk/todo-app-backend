package main

import (
	"fmt"
	"log"
	"todo/server"
)

func main() {
	svr := server.NewServer()
	err := svr.StartServer(3000)
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}

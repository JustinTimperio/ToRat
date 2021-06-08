package main

import (
	"log"

	_ "github.com/dimiro1/banner/autoload"
	server "github.com/JustinTimperio/ToRat/torat_server"
)

func main() {
	go func() {
		if err := server.Start(); err != nil {
			log.Println(err)
		}
	}()

	go server.APIServer()

	server.Shell()
}

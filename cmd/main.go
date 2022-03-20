package main

import (
	"log"
	"sartorius/converter/src/handler"
	"sartorius/converter/src/server"
)

func main() {
	h := &handler.Handler{}
	srv := new(server.Server)
	if err := srv.Run("8000", h.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}

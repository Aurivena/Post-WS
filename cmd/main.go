package main

import (
	"log"
	"net/http"
	"post-ws/pkg/ws"
)

func main() {
	http.HandleFunc("/", ws.WsConnect)
	run()
}

func run() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("dont run server")
		return
	}
	log.Println("server start")
}

package utils

import (
	"log"
	"net"
	"net/http"
	"time"

	"server"
)

func handler(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("S H I T"))
}

func handleConnection(conn net.Conn) {
	message := []byte("S H I T")
	_, err := conn.Write(message)
	if err != nil {
		log.Println("cant write", err)
	}
	defer conn.Close()
	time.Sleep(30 * time.Second)
}

func mainGrount() {
	http.HandleFunc("/", handler)
	log.Println("Server started")
	err := http.ListenAndServe(server.ADDRESS, nil)
	if err != nil {
		log.Println("Http err", err)
	}
}

func mainFraud() {
	listener, err := net.Listen(server.NETWORK, server.ADDRESS)
	if err != nil {
		log.Println("Cant listen on", err)
	}
	defer listener.Close()
	log.Println("Start listening")
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("cant accept", err)
		}
		log.Println("Connectid to user")
		go handleConnection(connection)
	}

}

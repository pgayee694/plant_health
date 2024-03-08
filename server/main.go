package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go RunSignals(sigs, done)

	server := mqtt.NewServer(nil)

	tcp := listeners.NewTCP("t1", ":1883")

	err := server.AddListener(tcp, nil)

	if err != nil {
		log.Fatal(err)
	}

	go RunServer(server)

	<-done
}

func RunServer(s *mqtt.Server) {
	err := s.Serve()

	if err != nil {
		log.Fatal(err)
	}
}

func RunSignals(sigs chan os.Signal, done chan bool) {
	<-sigs
	done <- true
}

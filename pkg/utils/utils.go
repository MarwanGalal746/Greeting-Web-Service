package utils

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func GetPort() string {
	port := flag.String("port", "8080", "The server port")
	log.Println("extracting port number from command line arguments")
	flag.Parse()
	log.Println("port number is: " + *port)
	return *port
}

func HandleKillSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Application started")

	//receiving the signal when it's sent through the channel
	go func() {
		receivedSignal := <-signalCh
		log.Println(receivedSignal.String() + ` has been received.`)
		log.Println("Shutting down gracefully")
		os.Exit(0)
	}()
}

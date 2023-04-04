package utils

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func GetPort() string {
	port := flag.String("port", "8080", "The server port")
	flag.Parse()
	return *port
}

func HandleKillSignals() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Application started")

	//receiving the signal when it's sent through the channel
	go func() {
		<-signalCh
		fmt.Println("Shutting down gracefully")
		os.Exit(0)
	}()
}

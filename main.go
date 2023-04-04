package main

import (
	"Greeting-web-service/pkg/handlers"
	"Greeting-web-service/pkg/logging"
	"Greeting-web-service/pkg/utils"
)

var portNum string

func init() {

	// Set up logging
	logging.Logging()

	//handle kill signals
	utils.HandleKillSignals()

	// Set up port
	portNum = utils.GetPort()

}

func main() {
	//get the port
	// Start the server
	handlers.Start(portNum)

	//todo : 1-documentation
}

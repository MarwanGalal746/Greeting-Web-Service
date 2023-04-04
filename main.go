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

	// Set up port
	portNum = utils.GetPort()

	//handle kill signals
	utils.HandleKillSignals()

}

func main() {
	//get the port
	// Start the server
	handlers.Start(portNum)

	//todo: 1-add tests. 2-logging. 3-documentation
}

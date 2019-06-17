package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/twuillemin/kuboxy/internal/configuration"
	"github.com/twuillemin/kuboxy/internal/controller"
	"github.com/twuillemin/kuboxy/pkg/context"
)

func main() {

	fmt.Printf("Starting...\n")

	// Get the configuration
	config, err := configuration.GetConfiguration()
	if err != nil {
		panic(err.Error())
	}
	config.Print()

	// Then load the possible configuration
	err = context.LoadContexts(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Contexts configuration loaded\n")

	// Display the configuration
	contextNames := context.GetContextNames()
	for _, contextName := range contextNames {
		fmt.Printf(" * %s\n", contextName)
	}

	fmt.Printf("Starting service\n")

	// Create the REST server
	go createServer(
		fmt.Sprintf("%s:%d", config.Address, config.RestPort),
		config.CertificateFileName,
		config.PrivateKeyFileName,
		controller.RegisterControllers)

	// Create the WebSocket server
	go createServer(
		fmt.Sprintf("%s:%d", config.Address, config.WebSocketPort),
		config.CertificateFileName,
		config.PrivateKeyFileName,
		controller.RegisterEventWebSocketController)

	// Wait until the end of the world
	<-make(chan interface{})
}

func createServer(address string, certificateFileName string, privateKeyFileName string, controllerRegistration func(e *echo.Echo)) {

	// Create an Echo server with the basic middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register the controllers
	controllerRegistration(e)

	// Start the Server
	var err error
	if len(certificateFileName) > 0 && len(privateKeyFileName) > 0 {
		err = e.StartTLS(address, certificateFileName, privateKeyFileName)
	} else {
		err = e.Start(address)
	}

	// In case of error, quit
	if err != nil {
		log.Fatalf("Unable to start server due to error: \"%v\"", err)
	}
}

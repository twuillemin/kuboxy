package controller

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
	"github.com/twuillemin/kuboxy/pkg/types"
)

// CommandType is the type of command sent to the EventListener
type CommandType string

const (
	// AddSource is the command for adding a new source
	AddSource CommandType = "AddSource"
	// RemoveSource is the command for deleting a source
	RemoveSource CommandType = "RemoveSource"
	// RemoveAllSources is the command for deleting all sources
	RemoveAllSources CommandType = "RemoveAllSources"
)

// Command is a command received by the event controller
type Command struct {
	Command       CommandType      `json:"command"`
	ObjectType    types.ObjectType `json:"objectType,omitempty"`
	ContextName   string           `json:"contextName,omitempty"`
	NamespaceName string           `json:"namespaceName,omitempty"`
}

// The source of an event
type eventSource struct {
	ObjectType    types.ObjectType
	ContextName   string
	NamespaceName string
}

// The information related to a forwarder
type forwarderInformation struct {
	source               eventSource
	stopForwarderChannel chan struct{}
}

// getEventsByWebSocket create the persistent websocket between a client and the server
// @Summary Connect to a WebSocket for managing events (port 8081)
// @Description the websocket used for receiving the configuration and then return the requested events. Each event is a full object when created / updated / deleted
// @ID get-events-by-websocket
// @Tags Events
// @Produce text/plain
// @Param contextName path string true "the name of the configuration"
// @Success 200 {string} string
// @Failure 404 {object} echo.HTTPError
// @Router /api/v1/events/ [get]
func getEventsByWebSocket(c echo.Context) (err error) {
	websocket.Handler(func(ws *websocket.Conn) {

		defer func() {
			err = ws.Close()
		}()

		forwarders := make([]*forwarderInformation, 0)

		var stopFlag struct{}
		sendMessageChannel := make(chan interface{})
		stopSendingChannel := make(chan struct{})
		stopHandlerChannel := make(chan struct{})

		// Start a goroutine for receiving client
		go func() {
			// Wait client message indefinitely
			for {
				msg := ""
				err = websocket.Message.Receive(ws, &msg)
				if err == nil {
					forwarders = processCommand(c, msg, forwarders, sendMessageChannel)
				} else {
					if err.Error() != "EOF" {
						c.Logger().Error(err)
					}
					stopSendingChannel <- stopFlag
					break
				}
			}
		}()

		// Start a goroutine for sending to the client
		go func() {
			// Loop endlessly
		sendLoop:
			for {
				// Block until we receive from one of following events
				select {
				case message := <-sendMessageChannel:
					jsonMessage, _ := json.Marshal(message)
					err = websocket.Message.Send(ws, string(jsonMessage))
					if err != nil {
						if err.Error() != "EOF" {
							c.Logger().Error(err)
						}
						break sendLoop
					}
				case <-stopSendingChannel:
					break sendLoop
				}
			}

			stopHandlerChannel <- stopFlag
		}()

		// Wait for the sending routine to terminate
		<-stopHandlerChannel

		// Stop all forwarders if any remaining
		forwarders = stopAllForwarders(forwarders)

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

// Process a command
func processCommand(c echo.Context, message string, forwarders []*forwarderInformation, sendMessageChannel chan interface{}) []*forwarderInformation {

	c.Logger().Info("Processing message: %s\n", message)

	commandReceived := new(Command)
	err := json.Unmarshal([]byte(message), &commandReceived)
	if err != nil {
		c.Logger().Error(err)
		return forwarders
	}

	source := eventSource{
		commandReceived.ObjectType,
		commandReceived.ContextName,
		commandReceived.NamespaceName,
	}

	switch commandReceived.Command {

	case AddSource:
		// Ensure previous provider does not exist
		var idxForwarder = getExistingProviderIndex(forwarders, source)
		if idxForwarder == -1 {
			forwarders, err = addEventSource(source, forwarders, sendMessageChannel)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			c.Logger().Warn(fmt.Sprintf("unable to add the source %v is already registered", source))
		}

	case RemoveSource:
		// Search previous forwarder
		var idxForwarder = getExistingProviderIndex(forwarders, source)
		// If previous forwarder found
		if idxForwarder > -1 {
			var stopFlag struct{}
			// Stop the forwarder
			forwarders[idxForwarder].stopForwarderChannel <- stopFlag
			// Remove from list (don't preserve the order)
			forwarders[idxForwarder] = forwarders[len(forwarders)-1]
			forwarders = forwarders[:len(forwarders)-1]
		}

	case RemoveAllSources:

		forwarders = stopAllForwarders(forwarders)
	}

	return forwarders
}

// stopAllForwarders stops all forwarders if any
func stopAllForwarders(forwarders []*forwarderInformation) []*forwarderInformation {

	var stopFlag struct{}
	for _, forwarder := range forwarders {
		// Stop the forwarder
		forwarder.stopForwarderChannel <- stopFlag
	}
	return make([]*forwarderInformation, 0)
}

// stopAllForwarders stops all forwarders if any
func getExistingProviderIndex(forwarders []*forwarderInformation, source eventSource) int {

	// Search previous forwarder
	for i := 0; i < len(forwarders); i++ {
		if forwarders[i].source == source {
			return i
		}
	}
	return -1
}

package handler

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		handlerMessage(v)
	case *events.Disconnected:
		fmt.Println("Disconnected")
	}
}

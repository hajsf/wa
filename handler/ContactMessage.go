package handler

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func ContactMessage(evt *events.Message) {
	Contact := evt.Message.GetContactMessage()
	fmt.Println(Contact.GetDisplayName())
	fmt.Println(Contact.GetVcard())
}

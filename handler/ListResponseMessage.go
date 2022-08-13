package handler

import (
	"go.mau.fi/whatsmeow/types/events"
)

func ListResponseMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	_, _ = sender, pushName
}

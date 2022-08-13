package handler

import (
	"go.mau.fi/whatsmeow/types/events"
)

func ButtonsResponseMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	_, _ = sender, pushName
}

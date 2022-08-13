package handler

import (
	"go.mau.fi/whatsmeow/types/events"
)

func handlerMessage(evt *events.Message) {
	if evt.Info.Chat.User == "status" {
	} else {
		switch {
		case evt.Message.Conversation != nil:
			go Conversation(evt)
		case evt.Message.ExtendedTextMessage != nil:
			go ExtendedTextMessage(evt)
		case evt.Message.DeviceSentMessage != nil:
			go DeviceSentMessage(evt)
		case evt.Message.Chat != nil:
			go Chat(evt)
		case evt.Message.ImageMessage != nil:
			go ImageMessage(evt)
		case evt.Message.StickerMessage != nil:
			go StickerMessage(evt)
		case evt.Message.AudioMessage != nil:
			go AudioMessage(evt)
		case evt.Message.VideoMessage != nil:
			go VideoMessage(evt)
		case evt.Message.DocumentMessage != nil:
			go DocumentMessage(evt)
		case evt.Message.ButtonsResponseMessage != nil:
			go ButtonsResponseMessage(evt)
		case evt.Message.ListResponseMessage != nil:
			go ListResponseMessage(evt)
		case evt.Message.LocationMessage != nil:
			go LocationMessage(evt)
		case evt.Message.ContactMessage != nil:
			go ContactMessage(evt)
		}
	}
}

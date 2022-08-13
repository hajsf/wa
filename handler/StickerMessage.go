package handler

import (
	"fmt"

	"go.mau.fi/whatsmeow/types/events"
)

func StickerMessage(evt *events.Message) {
	sticker := evt.Message.GetStickerMessage()
	if sticker.Url != nil {
		fmt.Println(sticker)
		/*	data, err := global.Cli.Download(audio)
			if err != nil {
				log.Printf("Failed to download audio: %v", err)
				return
			} */
		//	global.Passer.Logs <- fmt.Sprintf("Sticker: <a href='%v' target='_blank'>Open</a>", sticker.GetUrl())
	}
}

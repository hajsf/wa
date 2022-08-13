package handler

import (
	"fmt"
	"log"
	"mime"
	"os"
	"wa/api"
	"wa/utils"

	"go.mau.fi/whatsmeow/types/events"
)

func ImageMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	img := evt.Message.GetImageMessage()
	var caption string = ""
	if evt.Message.ImageMessage.Caption != nil {
		caption = *evt.Message.ImageMessage.Caption
	}

	if img != nil {
		file, err := api.Client.Download(img)
		if err != nil {
			log.Printf("Failed to download image: %v", err)
			return
		}
		exts, _ := mime.ExtensionsByType(img.GetMimetype())
		dirName := "waTemp"
		tmpDir := os.TempDir() + "\\" + dirName

		fileName := fmt.Sprintf("%s\\%s-%s%s", tmpDir, sender, evt.Info.ID, exts[0])
		path := fmt.Sprintf("http://localhost:1235/tmp/%s-%s%s", sender, evt.Info.ID, exts[0])

		if err = os.WriteFile(fileName, file, 0600); err != nil {
			log.Printf("Failed to save to server directory: %v", err)
		} else {
			log.Printf("Saved to server path at: %s", fileName)
		}
		data, _ := utils.PrepareModel(evt.Info.Chat.User,
			sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
			evt.Info.ID, "image", caption, "", path)
		api.Passer.Data <- api.SSEData{
			Event:   "message", // default: source.onmessage = function (event) {}
			Message: data,
		}
		//	global.Passer.Logs <- fmt.Sprintf("Image: <a href='%v' target='_blank'>Open</a>", path)
	}
}

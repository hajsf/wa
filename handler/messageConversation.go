package handler

import (
	"fmt"
	"wa/api"
	"wa/helpers"
	"wa/utils"

	"github.com/abadojack/whatlanggo"
	"go.mau.fi/whatsmeow/types/events"
)

func Conversation(evt *events.Message) {
	var msgReceived string

	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName
	received := evt.Message.GetConversation()
	// convert numbers in Arabic scrtip to numbers in latin script
	for _, e := range received {
		if e >= 48 && e <= 57 {
			//	fmt.Println("Number in english script number")
			msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
		} else if e >= 1632 && e <= 1641 {
			//	fmt.Println("It is Arabic script")
			msgReceived = fmt.Sprintf("%s%v", msgReceived, helpers.NormalizeNumber(e))
		} else {
			//	fmt.Println("Dose not looks to be a number")
			msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
		}
	}
	info := whatlanggo.Detect(evt.Message.GetConversation())
	fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence: ", info.Confidence)

	//	name := "Hasan"
	var lang string
	switch whatlanggo.Scripts[info.Script] {
	case "Arabic":
		//	go WelcomeMessage(sender, pushName)
		lang = "ar"
	case "Latin":
		lang = "en"
		//	go WelcomeMessageLatin(sender, pushName)
	}
	_ = lang
	// _ = i18n.HelloPerson(lang, name)

	/*	fmt.Println("Received a message!", evt.Message.GetConversation()) */

	data, _ := utils.PrepareModel(evt.Info.Chat.User,
		sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
		evt.Info.ID, "text", msgReceived, "", "")

	api.Passer.Data <- api.SSEData{
		Event:   "message", // default: source.onmessage = function (event) {}
		Message: data,
	}
}

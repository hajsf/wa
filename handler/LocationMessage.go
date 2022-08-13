package handler

import (
	"fmt"
	"wa/api"
	"wa/utils"

	"go.mau.fi/whatsmeow/types/events"
)

func LocationMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	Location := evt.Message.GetLocationMessage()
	fmt.Println(Location.GetDegreesLatitude())
	fmt.Println(Location.GetDegreesLongitude())
	fmt.Println(Location.GetAddress())

	latitude := Location.GetDegreesLatitude()
	longitud := Location.GetDegreesLongitude()
	_ = Location.GetAddress()
	link := fmt.Sprintf("<a href='https://www.google.com/maps/@%f,%f,15z' target='_blank'>Open map</a>", latitude, longitud)

	data, _ := utils.PrepareModel(evt.Info.Chat.User,
		sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
		evt.Info.ID, "location", "Location: "+link, "", "")
	api.Passer.Data <- api.SSEData{
		Event:   "message", // default: source.onmessage = function (event) {}
		Message: data,
	}

}

package utils

import (
	"encoding/json"
	"fmt"
)

type Model struct {
	Group, Sender, Name, Time, MessageID, MessageType, MessageText, MessageCaption, Uri string
}

func PrepareModel(Group, Sender, Name, Time, MessageID, MessageType, MessageText, MessageCaption, Uri string) (string, error) {
	model := Model{
		Group,
		Sender,
		Name,
		Time,
		MessageID,
		MessageType,
		MessageText,
		MessageCaption,
		Uri,
	}

	data, err := json.Marshal(model)
	if err != nil {
		fmt.Println(err)
		return "failed to JSON", err
	}
	//	fmt.Printf("Model: %#v", model)
	return string(data), nil

}

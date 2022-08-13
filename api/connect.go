package api

import (
	"context"
	"fmt"

	"go.mau.fi/whatsmeow"
)

var Client *whatsmeow.Client

func (p DataPasser) Connect() {
	if Client.IsConnected() {
		p.Data <- SSEData{
			Event:   "notification",
			Message: "Reconnecting to WhatsApp server ...",
		}
		Client.Disconnect()
	}

	if Client.Store.ID == nil {
		// No ID stored, new login
	GetQR:
		qrChan, _ := Client.GetQRChannel(context.Background())
		err := Client.Connect()
		if err != nil {
			//	panic(err)
			p.Data <- SSEData{
				Event:   "notification",
				Message: "Can not connect with WhatApp server, try again later",
			}
			fmt.Println("Sorry", err)

		}

		for evt := range qrChan {
			switch evt.Event {
			case "success":
				{
					p.Data <- SSEData{
						Event:   "granted",
						Message: "success",
					}
					fmt.Println("Login event: success")
				}
			case "timeout":
				{
					p.Data <- SSEData{
						Event:   "notification",
						Message: "Timeout or error reading from WhatsApp websocket, trying refreshing ...",
					}
					fmt.Println("Login event: timeout")
					goto GetQR
				}
			case "code":
				{
					fmt.Println("new code recieved")
					fmt.Println(evt.Code)
					p.Data <- SSEData{
						Event:   "qrCode",
						Message: evt.Code,
					}
				}
			}
		}
	} else {
		// Already logged in, just connect
		p.Data <- SSEData{
			Event:   "granted",
			Message: "Already logged in",
		}
		fmt.Println("Already logged")
		err := Client.Connect()
		if err != nil {
			fmt.Println("failed to connect to client:", err)
			// panic(err)
		}
	}
}

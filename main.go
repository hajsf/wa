package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"wa/api"
	"wa/global"
	"wa/server"
	"wa/ui"
)

func init() {
	// create a context that we can cancel
	global.Ctx = context.Background()
	global.Ctx, global.Cancel = context.WithCancel(global.Ctx)
}

func main() {
	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}

	wg.Add(1)
	go ui.Init(global.Ctx, &wg)

	go func() {
		server.Run(global.Ctx, ":1235")
		println("server closed")
	}()

	println("server started")

	// listen for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case <-global.Ctx.Done():
		fmt.Println("Context had been cancelled")
		global.Cancel()
	case <-c:
		fmt.Println("Server aborted")

		// destroy the ui if it is opened
		if ui.Wv != nil {
			fmt.Println("Destroying the ui")
			wg.Done()
			ui.Wv.Destroy()
		}

		fmt.Println("Destroying the wa client")
		if api.Client.IsConnected() {
			api.Passer.Data <- api.SSEData{
				Event:   "notification",
				Message: "Server is shut down at the host machine...",
			}
			api.Client.Disconnect()
		}
	}

	// and wait for them both to reply back
	wg.Wait()
	fmt.Println("main: all goroutines have told us they've finished")
}

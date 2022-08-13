package ui

import (
	"context"
	"log"
	"os"
	"sync"
	"wa/global"
	"wa/screen"

	"github.com/jchv/go-webview2"
) // go install github.com/jchv/go-webview2@latest

var Wv webview2.WebView

func Init(ctx context.Context, wg *sync.WaitGroup) {
	// tell the caller we've stopped
	defer wg.Done()

	width, height := screen.GetSystemMetrics(screen.SM_CXSCREEN), screen.GetSystemMetrics(screen.SM_CYSCREEN)

	Wv = webview2.NewWithOptions(webview2.WebViewOptions{
		Window:    nil,
		Debug:     false,
		DataPath:  "",
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "WhatsApp client",
			Width:  uint(width),
			Height: uint(height),
			IconId: 2,
			Center: true},
	})
	if Wv == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer Wv.Destroy()

	// Wv.SetTitle("WhatsApp CRM")
	Wv.SetSize(width, height, webview2.HintMax) // .HintFixed

	// load a local HTML file.
	_, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// w.Navigate(filepath.Join(c, "templates/index.html"))
	Wv.Navigate("http://localhost:1235/scan/")

	// bind JavaScript functions
	JSbinding()

	// Wv.Init("alert('hi')")

	Wv.Run()
	println("ui closed")
	global.Cancel()
}

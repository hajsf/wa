package waSocket

import (
	"fmt"
	"wa/api"
	"wa/global"
	"wa/handler"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func WaConnection(appName string) {
	//	var Client *whatsmeow.Client
	store.DeviceProps.Os = proto.String(appName)
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:datastore.db?_foreign_keys=on", dbLog)
	if err != nil {
		fmt.Println("error opening database", err)
		// panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		fmt.Println("error connecting device:", err)
		// panic(err)
	}
	//clientLog := waLog.Stdout("Client", "DEBUG", true)
	// Let's use our customLogger
	global.Log = LogText("Client", "DEBUG", true)

	api.Client = whatsmeow.NewClient(deviceStore, global.Log)
	api.Client.AddEventHandler(handler.EventHandler)
}

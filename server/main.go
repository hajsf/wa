package server

import (
	"context"
	"net/http"
	"wa/api"
	"wa/public"
	"wa/waSocket"
)

func Run(ctx context.Context, port string) {

	// launch the WhatsApp client
	waSocket.WaConnection("CRM")
	/*	if fsys, err := fs.Sub(public.Views, "../public"); err != nil {
			fmt.Println("error in reading embeded files:", err)
		} else {
			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))

			//http.Handle("/static/", http.FileServer(http.FS(fsys)))
		} */

	tmpDir := createTemp()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(public.Views))))
	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir(tmpDir))))

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/scan/", scan)
	http.HandleFunc("/messages/", messages)
	http.HandleFunc("/reply/", reply)
	http.HandleFunc("/jsonRes/", jsonRes)

	http.HandleFunc("/sse", api.Passer.HandleSignal)
	http.ListenAndServe(port, nil)

}

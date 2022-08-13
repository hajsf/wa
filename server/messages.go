package server

import (
	"fmt"
	"html/template"
	"net/http"
	"wa/public"
)

func messages(w http.ResponseWriter, r *http.Request) {

	if tmpl, err := template.ParseFS(public.Views,
		"layouts/base.html",
		"views/messages.html"); err != nil {
		fmt.Println("Error in file parsing:", err)
	} else {
		err = tmpl.ExecuteTemplate(w, "messages.html", nil)
		if err != nil {
			fmt.Println("error executing the template:", err)
		}
	}
}

package server

import (
	"fmt"
	"html/template"
	"net/http"
	"wa/public"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func index(w http.ResponseWriter, r *http.Request) {
	//	view = template.Must(template.ParseFS(globals.Views, "templates/layouts/*.html", "templates/views/*.html", "templates/partials/*.html"))
	//	if err != nil {
	//		log.Fatal("Error loading templates:" + err.Error())
	//	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	if tmpl, err := template.ParseFS(public.Views,
		"layouts/base.html",
		"views/home.html"); err != nil {
		fmt.Println("Error in file parsing:", err)
	} else {
		/*	err = tmpl.Execute(w, data)
			if err != nil {
				fmt.Println("error executing the template:", err)
			} */
		err = tmpl.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			fmt.Println("error executing the template:", err)
		}
	}
}

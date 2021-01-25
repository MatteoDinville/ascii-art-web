package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	ascii "./ascii"
)

type Data struct {
	Value string
}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/respons.html", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		Value := r.FormValue("ascii")
		namefile := r.FormValue("namefileed")
		Aprint := Data{Value: ascii.Ascii(Value, namefile)}
		templates, _ := template.ParseFiles("./static/respons.html")
		templates.ExecuteTemplate(w, "result", Aprint)

	})
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./static/index.html")
		tmpl.Execute(w, nil)
	})

	fmt.Printf("Server is listening to port #8090 ... \n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

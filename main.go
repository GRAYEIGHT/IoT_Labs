package main

import (
	"appengine"
    "net/http"
    "html/template"
)

func init() {
    http.HandleFunc("/", mainHandler)
}

var mainTemplate = template.Must(template.ParseFiles("index.html"))

func mainHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	err := mainTemplate.Execute(w, "")
	if err != nil {
		c.Errorf("mainTemplate: %v", err)
	}
}
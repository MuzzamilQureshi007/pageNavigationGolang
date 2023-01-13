package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/MuzzamilQureshi007/multiPageNav/pkg/config"
	handler "github.com/MuzzamilQureshi007/multiPageNav/pkg/handlers"
	render "github.com/MuzzamilQureshi007/multiPageNav/pkg/render"
)

const portNumber = ":7007"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/services", handler.Services)
	http.HandleFunc("/blog", handler.Blog)
	http.HandleFunc("/contact", handler.Contact)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/signup", handler.Signup)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.HandleFunc("/", handler.Static)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// http.Handle("/css/", http.FileServer(http.Dir("../static")))

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

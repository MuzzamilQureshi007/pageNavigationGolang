package handler

import (
	"net/http"

	"github.com/MuzzamilQureshi007/multiPageNav/pkg/render"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

// func Static(w http.ResponseWriter, r *http.Request) {
// 	http.FileServer(http.Dir("css/"))

// }

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

// About is the handler for the about page
func Services(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "services.page.tmpl")
}

// blog is the handler for the about page
func Blog(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "blog.page.tmpl")
}

// blog is the handler for the about page
func Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl")
}

// Login is the handler for the about page
func Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl")
}

// signup is the handler for the about page
func Signup(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "signup.page.tmpl")
}

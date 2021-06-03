package main

import (
	"fmt"
	"log"
	"net/http"

	"golang_practice/lenslocked.com/Working_project/views"

	"github.com/gorilla/mux"
)

var (
	HomeView    *views.View //stores the template and executes the template like html templates.
	ContactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// w.Header().Set("Content-Type", "text/plain")
	// err := HomeView.Template.ExecuteTemplate(w, HomeView.Layout, nil)
	//err := HomeView.Template.Execute(w, nil) //Executes The template
	must(HomeView.Render(w, nil))
	fmt.Println("Someone visited home page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// err := ContactView.Template.ExecuteTemplate(w, ContactView.Layout, nil) //calling a particular layout
	//err := ContactView.Template.Execute(w, nil) //Executes The template
	must(ContactView.Render(w, nil))
	fmt.Println("Someone visited contact page")

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	defer fmt.Println("Closed")
	// router := &http.ServeMux{}
	// router.HandleFunc("/", handlerFunc)
	//creating new router
	HomeView = views.NewView("bootstrap", "views/home.gohtml") //created viewfunction to handle the views
	ContactView = views.NewView("bootstrap", "views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	log.Fatalln(http.ListenAndServe(":8000", r))
}

//Adding new pages to our application and we will be using third package router which is gorilla/mux

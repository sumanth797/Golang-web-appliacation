package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	HomeTemplate    *template.Template //stores the template and executes the template like html templates.
	ContactTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// w.Header().Set("Content-Type", "text/plain")
	if err := HomeTemplate.Execute(w, nil); err != nil { //Executes The template
		panic(err)
	}
	fmt.Println("Someone visited home page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := ContactTemplate.Execute(w, nil); err != nil { //Executes The template
		panic(err)
	}
	fmt.Println("Someone visited contact page")

}
func main() {
	defer fmt.Println("Closed")
	// router := &http.ServeMux{}
	// router.HandleFunc("/", handlerFunc)
	//creating new router
	var err error
	HomeTemplate, err = template.ParseFiles("views/home.gohtml", "views/layout/footer.gohtml") //Execute the template in main function before our server starts
	if err != nil {
		panic(err)
	}
	ContactTemplate, err = template.ParseFiles("views/contact.gohtml", "views/layout/footer.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	log.Fatalln(http.ListenAndServe(":8000", r))
}

//Adding new pages to our application and we will be using third package router which is gorilla/mux

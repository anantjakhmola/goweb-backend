package main

import (
	"fmt"
	"net/http"

	"github.com/anantjakhmola/goweb-backend/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type,", "text/html")
	//fmt.Fprint(w, "<h1>Welcomne to my awesome site</h1>")
	must(homeView.Render(w, nil))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type,", "text/html")
	//fmt.Fprintf(w, "To get in touch please send email to <a href=\"mailto:anantjakhmola9@gmail.com\">support@anantmail.com</a>")
	must(contactView.Render(w, nil))
}
func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type,", "text/html")
	fmt.Fprintf(w, "<h1> Hey this is the FAQ page</h1>")
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

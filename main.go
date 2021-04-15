package main

import (
	"net/http"

	"github.com/anantjakhmola/goweb-backend/controllers"
	"github.com/anantjakhmola/goweb-backend/views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
)

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

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
	must(faqView.Render(w, nil))
	//fmt.Fprintf(w, "<h1> Hey this is the FAQ page</h1>")
}

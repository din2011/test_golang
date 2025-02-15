package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email "+
			"to <a href=\"mailto:support@lenslocked.com\">"+
			"support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you "+
			"were looking for :(</h1>"+
			"<p>Please email us if you keep being sent to an "+
			"invalid page.</p>")
	}
}

func main() {
	//http.HandleFunc("/", handlerFunc)
	//http.ListenAndServe(":3000", nil)

	var h http.Handler = http.HandlerFunc(notfound_router)
	
	

	r := mux.NewRouter()
	r.HandleFunc("/faq", faq_router)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	// This will assign the home page to the
	// NotFoundHandler
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func faq_router(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Page</h1>")
}

func notfound_router(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Page Not Found</h1>")
}


func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

package main

import (
	"net/http"
	route "tst/functions"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", route.Home)
	http.HandleFunc("/ascii", route.PrintAscii)
	http.ListenAndServe(":9000", nil)
}

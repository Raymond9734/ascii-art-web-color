package route

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	Ascii "tst/ascii-generator/banner"
)

func PrintAscii(w http.ResponseWriter, r *http.Request) {

	if strings.ToLower(r.Method) != "post" {
		http.Error(w, "Methhod not allowed", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/ascii" {
		http.Error(w, "Methhod not allowed", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("input")
	banner := r.FormValue("banner")

	str := strings.Split(text, "\r\n")
	var art string
	for _, wrd := range str {
		tempArt, err1 := Ascii.PrintBanner(wrd, banner)
		if err1 != nil {
			fmt.Println(err1)
			http.Error(w, "Internal server Error", http.StatusInternalServerError)
			return
		}
		art += tempArt
	}

	tmp, err := template.ParseFiles("templates/Home.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	contents := content{Ascii: art}
	tmp.Execute(w, contents)

}

package route

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("templates/Home.html")
	if err != nil {
		fmt.Println("homeError")
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}

	tmp.Execute(w, nil)

}

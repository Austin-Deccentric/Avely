package controller

import (
	"net/http"
	"log"
	"errors"
)

//HealthGet renders the accounts page
func healthGet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	err := tmpl.ExecuteTemplate(w, "healthindex.html", nil); if err != nil{
		log.Println("error loading template",err)
		http.Error(w, errors.New("Page is under construction").Error(), http.StatusNotFound)
	}
}
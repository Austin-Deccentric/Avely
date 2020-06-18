package controller

import (
	"errors"
	"log"
	"net/http"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	
	// Get the session from the store 
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Gets the user who has the session
	user := getUser(session)

	// Passing the user to our template
	err = tmpl.ExecuteTemplate(w, "dashboardindex.html", user); if err != nil{
		http.Error(w, errors.New("Something went wrong. If this continues contact an admin").Error(), http.StatusInternalServerError)
		log.Println("error loading template",err)
	}
	//tmpl.Execute(w, nil)
} 
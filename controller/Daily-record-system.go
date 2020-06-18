package controller

import (
	"log"
	"net/http"
	"errors"
)

func form(w http.ResponseWriter, r *http.Request) {
	// session, err := store.Get(r, "avely")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	log.Println(err)
	// 	return
	// }

	// user := getUser(session)

	// if auth := user.Authenticated; !auth {
	// 	session.AddFlash("You don't have access!")
	// 	err = session.Save(r, w)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		log.Println(err)
	// 		return
	// 	}
	// 	http.Redirect(w, r, "/forbidden", http.StatusFound)
	// 	return
	// }

	w.Header().Set("Content-Type", "text/html; charset=utf8")
	//tmpl, _ :=template.ParseFiles("./templates/sign_in.html")
	err = tmpl.ExecuteTemplate(w, "record-form.html", nil); if err != nil{
		http.Error(w, errors.New("Something went wrong. If this continues contact an admin").Error(), http.StatusInternalServerError)
		log.Println("error loading tmplate",err)
	}
	//tmpl.Execute(w, nil)
} 

func drsIndex(w http.ResponseWriter, r *http.Request) {
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
	err = tmpl.ExecuteTemplate(w, "record-page1.html", user); if err != nil{
		http.Error(w, errors.New("Something went wrong. If this continues contact an admin").Error(), http.StatusInternalServerError)
		log.Println("error loading tmplate",err)
	}
	//tmpl.Execute(w, nil)
} 



func secondPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	//tmpl, _ :=template.ParseFiles("./templates/sign_in.html")
	err := tmpl.ExecuteTemplate(w, "record-page2.html", nil); if err != nil{
		http.Error(w, errors.New("Something went wrong. If this continues contact an admin").Error(), http.StatusInternalServerError)
		log.Println("error loading template",err)
	}
	//tmpl.Execute(w, nil)
} 

func thirdPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	//tmpl, _ :=template.ParseFiles("./templates/sign_in.html")
	err := tmpl.ExecuteTemplate(w, "record-page3.html", nil); if err != nil{
		http.Error(w, errors.New("Something went wrong. If this continues contact an admin").Error(), http.StatusInternalServerError)
		log.Println("error loading tmplate",err)
	}
	//tmpl.Execute(w, nil)
} 


// func Drs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf8")
// 	//tmpl, _ :=template.ParseFiles("./templates/sign_in.html")
// 	http.Redirect(w, r, "https://ninyhorlah.github.io/Sensei/", http.StatusSeeOther)
// 	//tmpl.Execute(w, nil)
// } 
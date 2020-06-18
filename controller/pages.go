package controller

import (
	"fmt"
	"github.com/Austin-Deccentric/auth/views"
	"errors"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
)


//Signin renders the signin page
func signin(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, errors.New("Oops something went wrong. If this continues please contact the admin").Error(), http.StatusInternalServerError)
		return
	}
	user := getUser(session)

	if auth := user.Authenticated;  auth {
		// redirect to dashboard and return
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
		}
	
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	//tmpl, _ :=template.ParseFiles("./templates/sign_in.html")
	err = tmpl.ExecuteTemplate(w, "sign_in.html", nil); if err != nil{
		http.Error(w, errors.New("Page is under construction").Error(), http.StatusNotFound)
		log.Println("error loading tmplate",err)
	}
	//tmpl.Execute(w, nil)
} 

// // Signin renders the signin page
// func Signin(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("hello word")
// 	//w.Write([]byte("Hello world"))
// 	w.Header().Set("Content-Type", "text/html; charset=utf8")
// 	http.Redirect(w, r, "/static/sign_in.html", http.StatusSeeOther)
	
// 	//tmpl.Execute(w, nil)
// } 

// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) views.User {
	// Gets the session info
	val := s.Values["user"]
	// // if no session cookie
	// if !ok {
	// 	return views.User{Authenticated: false}
	// }

	var user = views.User{}
	// Checks if there is a user stored in the cookie
	user, ok := val.(views.User)
	if !ok {
		return views.User{Authenticated: false}
	}
	return user
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "forbidden.gohtml", session.Flashes())
}

// Secret renders a hidden message chore: To be removed
func secret(w http.ResponseWriter, r *http.Request) {
	// session, err := store.Get(r, "avely")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//user := getUser(session)
	fmt.Println("secrent function runs")
	tpl.ExecuteTemplate(w, "secret.gohtml", nil)
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "avely")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	// Checks if the cookie is new. If it is new, the user is not logged in
	if session.IsNew {
		session.AddFlash("You must be logged in to access this page")
		//http.Error(w, err.Error(), http.StatusUnauthorized)
		err = errors.New("User is not logged in")
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		tpl.ExecuteTemplate(w, "forbidden.gohtml", session.Flashes())

		return
	}
	
	user := getUser(session)

	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}
		w.WriteHeader(http.StatusUnauthorized)
		tpl.ExecuteTemplate(w, "forbidden.gohtml", session.Flashes())
		return
	}
	next(w,r)
	}
	
}
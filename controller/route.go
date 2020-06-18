package controller

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Register is where our routes and handlers live.
func Register() *mux.Router{
	r := mux.NewRouter().StrictSlash(true)
	//r.Handle("/", SigninPage).Methods("GET")
	//r.HandleFunc("/", controller.Index)
	r.HandleFunc("/", signin).Methods("GET") // root address renders signin page
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/signup", signup).Methods("GET", "POST")
	r.HandleFunc("/symptoms",authenticate(healthGet)).Methods("GET")//renders the symptoms (healthindex.html) page
	r.HandleFunc("/accounts", authenticate(accountsGet)).Methods("GET") //renders accounts html tamplate on server
	//r.HandleFunc("/refresh", controller.Refresh)
	r.HandleFunc("/forbidden", forbidden)
	r.Handle("/secret", authenticate(secret))

	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", js))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", fimg))

	// mux for the daily record system
	r.HandleFunc("/form", authenticate(form))
	r.HandleFunc("/firstpage", authenticate(drsIndex))
	r.HandleFunc("/secondpage", authenticate(secondPage))
	r.HandleFunc("/thirdpage", authenticate(thirdPage))
	r.HandleFunc("/dashboard", authenticate(dashboard))
	return r

}
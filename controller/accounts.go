package controller

import (
	"log"
	"encoding/json"
	"net/http"

	//"github.com/startng/sensei-poultry-management/model"
)

//Accounts is sent to the respective endpoint
func Accounts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		accountsGet(w, r)
	case "POST":
		accountsPost(w, r)
	}
}

//AccountsGet renders the accounts page
func accountsGet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	err := tmpl.ExecuteTemplate(w, "accountingindex.html", nil); if err != nil{
		log.Println("error loading template",err)
		w.WriteHeader(http.StatusNotFound)
	}
}

//AccountsPost takes input ans sends it to the database
func accountsPost(w http.ResponseWriter, r* http.Request){
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"Page under construction"})

}
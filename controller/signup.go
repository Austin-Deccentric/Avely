package controller

import (
	"github.com/Austin-Deccentric/auth/model"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"github.com/badoux/checkmail"
)

const hashCost = 8


//Signup sWitches between requests get and post in SignupGet and SignupPost
func signup(w http.ResponseWriter, r*http.Request) {
	switch r.Method {
	case "GET":
		signupGet(w, r)
	case "POST":
		signupPost(w, r)
		
	}
}

//signupGet renders the signup ppage
func signupGet(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	err := tmpl.ExecuteTemplate(w, "index.html", nil); if err != nil{
		log.Println("error loading tmplate",err)
	}
}


// signupPost registers a new user in the database
func signupPost(w http.ResponseWriter, r *http.Request){
	// Parse and decode the request body into a new `Credentials` instance
	username :=  r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	//checks validity of the host and format  of the incoming email 
	err := checkmail.ValidateFormat(email)
	if err != nil {
		http.Error(w, errors.New("Invalid email address").Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	
	// err = checkmail.ValidateHost(email)
	// if err != nil {
	// 	http.Error(w, errors.New("Invalid email address").Error(), http.StatusBadRequest)
	// 	log.Println("Invalid email adress", err)
	// 	return
	// }

	// err = checkmail.ValidateHost(email)
	// 	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
	// 		http.Error(w, errors.New("Invalid email address").Error(), http.StatusBadRequest)
	// 		log.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
	// 		return
	// 	}
	
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	
	passkey := string(hashedPassword)
	// Next, insert the username, along with the email hashed , password into the database
	if err = model.CreateUser(username, email, passkey); err != nil {
		// If there is any issue with inserting into the database, return a 400 error
		http.Error(w, errors.New("User already registered").Error(), http.StatusBadRequest)
		log.Println("Duplicate entry",err)
		return
	}
	// redirect to signin
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	http.Redirect(w, r, "/", http.StatusSeeOther)

	//json.NewEncoder(w).Encode(creds.Username)
	// We reach this point if the credentials were correctly stored in the database, and the default status of 200 is sent back
}
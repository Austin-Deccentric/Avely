package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/Austin-Deccentric/auth/model"
	//"github.com/twinj/uuid"
	//"time"
	"encoding/gob"
	"github.com/Austin-Deccentric/auth/views"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	//"path/filepath" // so that we can make path joins compatible on all OS
	"github.com/badoux/checkmail"
	//redisStore "gopkg.in/boj/redistore.v1"
)


var (
	//store *redisStore.RediStore

	//Fs is serves static files on the server
	fs = http.FileServer(http.Dir("./templates/css/"))
	js = http.FileServer(http.Dir("./templates/js/"))
	fimg = http.FileServer(http.Dir("./templates/img/"))

	tmpl = template.New("")
	 
	tpl *template.Template
	err error
	store *sessions.CookieStore
)
func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)  // encryption key is optional

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	// _, err := tmpl.ParseGlob(filepath.Join(".", "templates", "*.html")) // parsed templates for use in pages.go
	// if err != nil {
	// 	log.Printf("Unable to parse templates: %v\n", err)
	// }

	//Implementation of a redis store to store cookies. Removed beacause i could not reproduce on Heroku. Todo: Retry
	//store,err =  redisStore.NewRediStore(5, "tcp", ":6379","",authKeyOne,encryptionKeyOne)
	//pleasenaawor :=  NewRediStore(5, "tcp", ":6379","",authKeyOne,encryptionKeyOne)
		// if err != nil {
		// 	log.Println(err)
		// }
		// fmt.Println("Connected to cache memory")
	//defer store.Close()                         

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		// Secure: true,
		HttpOnly: true,
	}

	gob.Register(views.User{})
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	files := []string {"templates/daily-record-system/record-form.html",
						"templates/daily-record-system/record-page1.html",
						"templates/daily-record-system/record-page2.html",
						"templates/daily-record-system/record-page3.html",
						"templates/login/sign_in.html",
						"templates/login/index.html",
						"templates/acc-health/accountingindex.html",
						"templates/acc-health/healthindex.html",
						"templates/dashboard/dashboardindex.html"}  // A verbose implemtation of pasring templates 
	
	tmpl, err = template.ParseFiles(files...)
	if err != nil{
		log.Println("Unable to parse files",err)
	}
}


// Login authenticates user login credentials
func login(w http.ResponseWriter, r *http.Request){
	// Parse and decode the form
	email := r.FormValue("email")
	password := r.FormValue("password")

	// create a new session for user  chore: better error handling log it properly for deployment
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error decoding session",err)
		return
	}

	// Validate user email
	err = checkmail.ValidateFormat(email)
		if err != nil {
			//session.AddFlash("Enter a valid email address")
			// try redirect back to login
			fmt.Fprint(w, "Enter a valid email address")
			log.Println("Invalid email format",err)
			return
		}
	

	
	// We create another instance of `Credentials` to store the credentials we get from the database
	hashedCreds,err := model.GetUserCredential(email)
	if err != nil {
		// If an entry with the email does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			log.Println("Email not in record",err)
			//w.WriteHeader(http.StatusUnauthorized)
			session.AddFlash("Email not found")
			w.WriteHeader(http.StatusUnauthorized)
			tpl.ExecuteTemplate(w, "forbidden.gohtml", session.Flashes())  // check if i get a superflous header error
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Somethin wrong with database",err)
		return
	}
	
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(hashedCreds.Password), []byte(password)); err != nil {
		// write to session, todo: will probaly use to keep count of attempts 
		err = session.Save(r, w)
		if err != nil {
			log.Println("Error saving session")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Error with authentication, incorrect password")
		// If the two passwords don't match, return a 401 status
		//w.WriteHeader(http.StatusUnauthorized)
		// Render error message on page
		session.AddFlash("The code was incorrect")
		w.WriteHeader(http.StatusUnauthorized)
		tpl.ExecuteTemplate(w, "forbidden.gohtml", session.Flashes())
		return
	}
	 

	user := views.User{
		Email:      email,
		Authenticated: true,
	}

	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	//w.WriteHeader(http.StatusOK)
	http.Redirect(w, r, "/dashboard", http.StatusFound)
	// If we reach this point, that means the users password was correct, and that they are authorized

}

// Logout revokes authentication for a user
func logout(w http.ResponseWriter, r *http.Request) {  // todo: make this only accessible to logged in users
	session, err := store.Get(r, "avely")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	session.Values["user"] = views.User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)  // redirect to login page
}


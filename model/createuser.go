package model

import (
	"database/sql"
	"github.com/Austin-Deccentric/auth/views"
)

// CreateUser creates a new user
func CreateUser(username, email, password string) error {
	// before creating user check if the user name exits
	err := checkduplicateuser(email); if err != nil { 
		return err
	}
	if _, err = db.Query("INSERT INTO users_info (username, email, password) VALUES ($1,$2,$3)", username, email, password); err != nil {
		return err
	}
	return nil
}

func checkduplicateuser(email string) error {
	result := db.QueryRow("SELECT email FROM users_info WHERE email= $1", email)
	userEmail := &views.Credentials{}
	// Store the obtained user in `storedCreds`
	err := result.Scan(&userEmail.Email)
	if err == sql.ErrNoRows {
		//log.Println("Error ocurred parsing password", err)
		return nil
	}
	//fmt.Println(err,userEmail)
	return err
}
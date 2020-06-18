package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/Austin-Deccentric/auth/controller"
	"github.com/Austin-Deccentric/auth/model"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	//"path/filepath" // so that we can make path joins compatible on all OS
)


func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port ==""{
		port = strconv.Itoa(8000)
	}
	  
	//SigninPage := http.FileServer(http.Dir("./templates/sign_in.html"))
	
	// initialize our Postgres database connection
	db := model.InitDB()
	defer db.Close()
	r := controller.Register()

	fmt.Printf("Listening and serving on port %s.....\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
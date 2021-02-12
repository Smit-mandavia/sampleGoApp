package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,  "Hello World")
}

func handleRequests(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/users", AllUsers).Methods("GET")
	//router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser ).Methods("POST")
	//router.HandleFunc("/users/{id}", UpdateUser).Methods("PATCH")
	//router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}


func main() {
	DatabaseInit()
	//defer DatabaseClose()
	DatabaseMigrate()
	handleRequests()
}
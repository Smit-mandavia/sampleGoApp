package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)


var err error
var db *gorm.DB
var dsn string = "host="+ DatabaseHost +" user=" + DatabaseUser + " password=" + DatabasePassword + " dbname=" + DatabaseName + " port=" + DatabasePort + " sslmode=" + DatabaseSslMode + " TimeZone=" + DatabaseTimeZone

type User struct {
	gorm.Model
	Name string
	Email string
}

func DatabaseInit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Database not connected")
		}
	}()

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println(err.Error)
		panic("failed database")
	}
}

func DatabaseMigrate(){
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request){
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}()

	nameInf := data["name"]
	emailInf := data["email"]

	name,ok1 := nameInf.(string)
	email,ok2 := emailInf.(string)

	if (ok1 && ok2) != true {
		http.Error(w, `{ "message" : "name and email field are required"`, http.StatusBadRequest)
		return
	}

	user := User{Name:name, Email:email}
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
	return
}

func GetUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "all")

}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "all")

}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "all")

}



//type event struct {
//	ID          string `json:"ID"`
//	Title       string `json:"Title"`
//	Description string `json:"Description"`
//}
//
//type allEvents []event
//
//var events = allEvents{
//	{
//		ID:          "1",
//		Title:       "Introduction to Golang",
//		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
//	var newEvent event
//	// Convert r.Body into a readable formart
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "Kindly enter data with the event id, title and description only in order to update")
//	}
//
//	json.Unmarshal(reqBody, &newEvent)
//
//	// Add the newly created event to the array of events
//	events = append(events, newEvent)
//
//	// Return the 201 created status code
//	w.WriteHeader(http.StatusCreated)
//	// Return the newly created event
//	json.NewEncoder(w).Encode(newEvent)
//}
//
//func getOneEvent(w http.ResponseWriter, r *http.Request) {
//	// Get the ID from the url
//	eventID := mux.Vars(r)["id"]
//
//	// Get the details from an existing event
//	// Use the blank identifier to avoid creating a value that will not be used
//	for _, singleEvent := range events {
//		if singleEvent.ID == eventID {
//			json.NewEncoder(w).Encode(singleEvent)
//		}
//	}
//}
//
//func getAllEvents(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(events)
//}
//
//func updateEvent(w http.ResponseWriter, r *http.Request) {
//	// Get the ID from the url
//	eventID := mux.Vars(r)["id"]
//	var updatedEvent event
//	// Convert r.Body into a readable formart
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
//	}
//
//	json.Unmarshal(reqBody, &updatedEvent)
//
//	for i, singleEvent := range events {
//		if singleEvent.ID == eventID {
//			singleEvent.Title = updatedEvent.Title
//			singleEvent.Description = updatedEvent.Description
//			events[i] = singleEvent
//			json.NewEncoder(w).Encode(singleEvent)
//		}
//	}
//}
//
//func deleteEvent(w http.ResponseWriter, r *http.Request) {
//	// Get the ID from the url
//	eventID := mux.Vars(r)["id"]
//
//	// Get the details from an existing event
//	// Use the blank identifier to avoid creating a value that will not be used
//	for i, singleEvent := range events {
//		if singleEvent.ID == eventID {
//			events = append(events[:i], events[i+1:]...)
//			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
//		}
//	}
//}//	},
////}
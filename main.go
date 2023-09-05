package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Food represent a single food option in the menu
type Food struct {
	Name           string              `json:"name"`
	Description    string              `json:"description"`
	Price          float64             `json:"price"`
	Especification *SpecialityOfTheDay `json:"especification"`
}

type SpecialityOfTheDay struct {
	Category  string `json:"category"`
	Available string `json:"available"`
}

var menu []Food //menu is a slice of Food that allows us to storage small amount of data without a database

// getToday retunrs the regular menu and the menu available at the current time
func getToday(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usertime := time.Now().UTC()                                                              //Setting up user real time whenever the function is called
	userToLisbonTime := usertime.In(time.FixedZone("Europe/Lisbon", 0000)).Weekday().String() //Setting the standart time to Lisbon

	var availableItems []Food //Storage only the items available in that weekday

	for _, item := range menu {
		if item.Especification.Available == userToLisbonTime || item.Especification.Available == "Everyday" {
			availableItems = append(availableItems, item)
		}
	}
	json.NewEncoder(w).Encode(availableItems)
}

// getAll returns the complete menu regardless the weekday
func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menu)
}

// getByDay returns the menu given a specific weekday
func getByDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "encoding/json")

	var availableByDay []Food

	params := mux.Vars(r) //entender o que faz essa variavel depois

	for _, item := range menu {
		if item.Especification.Available == params["available"] || item.Especification.Available == "Everyday" {
			availableByDay = append(availableByDay, item)
		}
	}
	json.NewEncoder(w).Encode(availableByDay)
}

func main() {
	//Inicializing the router
	router := mux.NewRouter()

	//Some examples
	menu = append(menu, Food{Name: "Alentejana", Description: "Diced pork, diced potatos and Clam", Price: 10.95, Especification: &SpecialityOfTheDay{Category: "Regular Menu", Available: "Everyday"}})
	menu = append(menu, Food{Name: "Strogonoff de Frango", Description: "Sliced chicken with cream sauce and mushrooms", Price: 10.95, Especification: &SpecialityOfTheDay{Category: "Speciality", Available: "Monday"}})
	menu = append(menu, Food{Name: "Panados de porco", Description: "Breaded pork served with white rice, fries and salad", Price: 10.95, Especification: &SpecialityOfTheDay{Category: "Speciality", Available: "Tuesday"}})

	//Setting up the route handlers
	router.HandleFunc("/menutoday", getToday).Methods("GET")
	router.HandleFunc("/menu", getAll).Methods("GET")
	router.HandleFunc("/menu/{available}", getByDay).Methods("GET")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}

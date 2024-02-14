package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type House struct {
	Hid      string `json:"hid"`
	HName    string `json:"hname"`
	HAddress string `json:"haddress"`
	HCity    string `json:"hcity"`
}

var houses []House

func getHouses(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

func getHouse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range houses {
		if item.Hid == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//json.NewEncoder(w).Encode(&House{})
}

func createHouse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newHouse House
	_ = json.NewDecoder(r.Body).Decode(&newHouse)
	newHouse.Hid = strconv.Itoa(len(houses) + 1)
	houses = append(houses, newHouse)
	json.NewEncoder(w).Encode(newHouse)
}

func updateHouse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range houses {
		if item.Hid == params["id"] {
			houses = append(houses[:i], houses[i+1:]...)
			var newHouse House
			_ = json.NewDecoder(r.Body).Decode(&newHouse)
			newHouse.Hid = params["id"]
			houses = append(houses, newHouse)
			json.NewEncoder(w).Encode(newHouse)
			return
		}
	}
}

func deleteHouse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range houses {
		if item.Hid == params["id"] {
			houses = append(houses[:i], houses[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(houses)
}
func main() {

	houses = append(houses, House{Hid: "1", HName: "Sri House", HAddress: "Vijaynagar", HCity: "Chennai"},
		House{Hid: "2", HName: "AP House", HAddress: "Vijaynagar", HCity: "Delhi"},
		House{Hid: "3", HName: "MI House", HAddress: "Vijaynagar", HCity: "Mumbai"})

	//endpoints
	router := mux.NewRouter()
	fmt.Println("Main Invoked")
	router.HandleFunc("/houses", getHouses).Methods("GET")
	router.HandleFunc("/houeses/{id}", getHouse).Methods("GET")
	router.HandleFunc("/houses", createHouse).Methods("POST")
	router.HandleFunc("/houses/{id}", updateHouse).Methods("PUT")
	router.HandleFunc("/houses/{id}", deleteHouse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":6001", router))

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Location struct {
	ID        uint32  `json:"id"`
	Distance  float64 `json:"distance"`
	PlaceName string  `json:"placeName"`
	Address   string  `json:"address"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	Ac        bool    `json:"ac"`
	Wifi      bool    `json:"wifi"`
}

func SetupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SetupResponse(&w, r)

	var location Location
	location.ID = 1
	location.Distance = 500.3
	location.PlaceName = "Bandung"
	location.Address = "Jl. Buah Batu"
	location.Lat = 110.2
	location.Long = 236.4
	location.Ac = true
	location.Wifi = true

	json.NewEncoder(w).Encode(location)
}

func main() {
	//START -- port//
	var portNet = os.Getenv("PORT")
	if portNet == "" {
		portNet = "8000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + portNet)
	}
	//END -- port//

	//START -- mux route -- //
	router := mux.NewRouter()

	router.HandleFunc("/location", getLocation).Methods("GET")

	http.ListenAndServe(":"+portNet, router)
	//END -- mux route -- //
}

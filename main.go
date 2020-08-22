package main

import (
	"github.com/GeistInDerSH/Covid19-Watch/covid_data"
	"github.com/GeistInDerSH/Covid19-Watch/csv_data"
	"github.com/gorilla/mux"

	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	MergedData []*covid.CombinedData
)

// returnAllCounties returns a json list containing all countries, their
// confirmed cases, deaths, and recoveries.
func returnAllCountries(w http.ResponseWriter, r *http.Request) {

	if MergedData == nil {
		updateCovidData()
	}

	json.NewEncoder(w).Encode(MergedData)
}

// returnSingleCountry checks for the given country matching the id.
// If it is found, the json data will be written to the page.
func returnSingleCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if MergedData == nil {
		updateCovidData()
	}

	// Use the 2 letter country code
	if len(key) == 2 {
		key = strings.ToUpper(key)
		for _, country := range MergedData {
			if country.Id == key {
				json.NewEncoder(w).Encode(country)
				break
			}
		}
	} else {
		for _, country := range MergedData {
			if country.Country == key {
				json.NewEncoder(w).Encode(country)
				break
			}
		}
	}
}

// handleRequests sets up the different web pages and the functions
// associated with those pages.
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/countries", returnAllCountries)
	router.HandleFunc("/country/{id}", returnSingleCountry)
	log.Fatal(http.ListenAndServe(":80", router))
}

// updateCovidData retreives the confirmed cases, deaths, and recovered
// cases for each of the countries. These three are then merged together.
func updateCovidData() {
	confirmedData, err := csvdata.GetCSVData("confirmed")
	if err != nil {
		log.Fatal("Could not download confirmed data")
	}

	deathData, err := csvdata.GetCSVData("deaths")
	if err != nil {
		log.Fatal("Could not download death data")
	}

	recoveredData, err := csvdata.GetCSVData("recovered")
	if err != nil {
		log.Fatal("Could not download recovered data")
	}

	MergedData = covid.MergeResults(confirmedData, deathData, recoveredData)
}

func main() {
	ticker := time.NewTicker(12 * time.Hour)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				updateCovidData()

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	handleRequests()
}

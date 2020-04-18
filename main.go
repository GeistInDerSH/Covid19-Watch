package main

import (
	"github.com/GeistInDerSH/Covid19-Watch/covid_data"
	"github.com/GeistInDerSH/Covid19-Watch/csv_data"

	"log"
	"sort"
)

// sortResults sorts the covid data based on the number of deaths from highest to lowest
func sortResults(covidData []*covid.CombinedCovidData) []*covid.CombinedCovidData {
	sort.Slice(covidData[:], func(i, j int) bool {
		return covidData[i].Deaths > covidData[j].Deaths
	})

	return covidData
}

func main() {
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

	mergedData, err := covid.MergeResults(confirmedData, deathData, recoveredData)
	if err != nil {
		log.Fatal("Could not merge the data")
	}

	mergedData = sortResults(mergedData)

	for _, d := range mergedData {
		d.PrintData()
	}
}

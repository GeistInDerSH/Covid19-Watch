package csvdata

import (
	"github.com/GeistInDerSH/Covid19-Watch/covid_data"

	"fmt"
	"net/http"
)

const JohnHopkinsURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_%s_global.csv"

// downloadCSVData attempts to download the CSV data from the given URL and store it into a map
func downloadCSVData(url string) (map[string]*covid.SingleData, error) {
	var data map[string]*covid.SingleData
	response, err := http.Get(url)
	if err != nil {
		return data, err
	}

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("Got %q when getting data from %q", response.StatusCode, url)
		return data, err
	}
	defer response.Body.Close()

	data, err = covid.ParseCSVData(response.Body)
	return data, err
}

// GetCSVData downloads the CSV data for the given sufix.
// The valid suffixes are: confirmed, deaths, or recovered
func GetCSVData(sufix string) (map[string]*covid.SingleData, error) {
	return downloadCSVData(fmt.Sprintf(JohnHopkinsURL, sufix))
}

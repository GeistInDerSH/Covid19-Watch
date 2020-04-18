package covid

import (
	"github.com/GeistInDerSH/Covid19-Watch/color"
	"github.com/olekukonko/tablewriter"

	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type CombinedCovidData struct {
	Region    string
	Date      string
	Infected  int
	Deaths    int
	Recovered int
}

type SingleCovidData struct {
	Region string
	Date   string
	Value  int
}

// PrintData prints the data stored in the CombinedCovidData object in a table
func (c *CombinedCovidData) PrintData() {
	date := color.SetColor(color.Purple, "Date:") + color.SetColor(color.None, c.Date)
	confirmedCases := color.SetColor(color.Purple, "Confirmed:") + color.SetColor(color.Blue, strconv.Itoa(c.Infected))
	deaths := color.SetColor(color.Purple, "Deaths:") + color.SetColor(color.Red, strconv.Itoa(c.Deaths))
	recoveredCases := color.SetColor(color.Purple, "Recovered:") + color.SetColor(color.Green, strconv.Itoa(c.Recovered))

	data := [][]string{
		[]string{"", date, ""},
		[]string{confirmedCases, deaths, recoveredCases},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetColWidth(40)
	table.SetAutoMergeCells(true)
	table.SetCenterSeparator("-")

	table.SetHeader([]string{"", c.Region, ""})
	table.SetHeaderColor(tablewriter.Colors{},
		tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{})

	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.AppendBulk(data)
	table.Render()
	fmt.Println()
}

// ParseCSVData parses the CSV data into a map with the country name as the key and the combined
// confirmed cases, total deaths, and recovered cases for each state within the country
func ParseCSVData(input io.Reader) (map[string]*SingleCovidData, error) {
	data := make(map[string]*SingleCovidData)
	csvReader := csv.NewReader(input)
	csvReader.TrimLeadingSpace = true

	headers, err := csvReader.Read()
	if err != nil {
		return data, err
	}

	csvReader.FieldsPerRecord = len(headers)
	mostRecentDate := headers[len(headers)-1]

	for {
		cell, err := csvReader.Read()
		if len(cell) == 0 {
			break
		}
		if err != nil {
			return data, err
		}

		mostRecentValue, err := strconv.Atoi(cell[len(cell)-1])
		if err != nil {
			break
		}

		country := cell[1]
		if _, exists := data[country]; exists {
			data[country].Value += mostRecentValue
		} else {
			data[country] = &SingleCovidData{
				Region: country,
				Date:   mostRecentDate,
				Value:  mostRecentValue,
			}
		}
	}

	return data, err
}

// stringCleanup removes any special characters that exists within the string and
// if the country is broken up into two parts it corrects it (e.g Korea, South -> South Korea)
func stringCleanup(s string) string {
	i := strings.Index(s, "*")
	if i > 0 {
		runeArray := []rune(s)
		s = string(runeArray[:i]) + string(runeArray[i+1:])
	}

	i = strings.Index(s, ", ")
	if i > 0 {
		runeArray := []rune(s)
		s = string(runeArray[i+2:]) + " " + string(runeArray[:i])
	}

	return s
}

// MergeResults merges the individual daata for the confirmed covid cases, deaths, & recovered cases
func MergeResults(confirmedCases, deaths, recoveredCases map[string]*SingleCovidData) ([]*CombinedCovidData, error) {
	var covidData []*CombinedCovidData
	for _, confirmed := range confirmedCases {
		deathsObj := deaths[confirmed.Region]
		recovredObj := recoveredCases[confirmed.Region]
		addCovidData := &CombinedCovidData{
			Region:    stringCleanup(confirmed.Region),
			Date:      confirmed.Date,
			Infected:  confirmed.Value,
			Deaths:    deathsObj.Value,
			Recovered: recovredObj.Value,
		}

		covidData = append(covidData, addCovidData)
	}

	return covidData, nil
}

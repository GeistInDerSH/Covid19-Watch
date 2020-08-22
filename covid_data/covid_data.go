package covid

import (
	"encoding/csv"
	"io"
	"sort"
	"strconv"
	"strings"
)

var countryIds = map[string]string{
	"Afghanistan":                       "AF",
	"Albania":                           "AL",
	"Algeria":                           "DZ",
	"American Samoa":                    "AS",
	"Andorra":                           "AD",
	"Angola":                            "AO",
	"Anguilla":                          "AI",
	"Antarctica":                        "AQ",
	"Antigua and Barbuda":               "AG",
	"Argentina":                         "AR",
	"Armenia":                           "AM",
	"Aruba":                             "AW",
	"Australia":                         "AU",
	"Austria":                           "AT",
	"Azerbaijan":                        "AZ",
	"Bahamas":                           "BS",
	"Bahrain":                           "BH",
	"Bangladesh":                        "BD",
	"Barbados":                          "BB",
	"Belarus":                           "BY",
	"Belgium":                           "BE",
	"Belize":                            "BZ",
	"Benin":                             "BJ",
	"Bermuda":                           "BM",
	"Bhutan":                            "BT",
	"Bolivia":                           "BO",
	"Bonaire":                           "BQ",
	"Bosnia and Herzegovina":            "BA",
	"Botswana":                          "BW",
	"Bouvet Island":                     "BV",
	"Brazil":                            "BR",
	"British Indian Ocean Territory":    "IO",
	"Brunei":                            "BN",
	"Bulgaria":                          "BG",
	"Burkina Faso":                      "BF",
	"Burundi":                           "BI",
	"Cabo Verde":                        "CV",
	"Cambodia":                          "KH",
	"Cameroon":                          "CM",
	"Canada":                            "CA",
	"Cape Verde":                        "CV",
	"Cayman Islands":                    "KY",
	"Central African Republic":          "CF",
	"Chad":                              "TD",
	"Chile":                             "CL",
	"China":                             "CN",
	"Christmas Island":                  "CX",
	"Cocos (Keeling) Islands":           "CC",
	"Colombia":                          "CO",
	"Comoros":                           "KM",
	"Congo (Brazzaville)":               "CG",
	"Congo (Kinshasa)":                  "CD",
	"Cook Islands":                      "CK",
	"Costa Rica":                        "CR",
	"Croatia":                           "HR",
	"Cuba":                              "CU",
	"Curacao":                           "CW",
	"Cyprus":                            "CY",
	"Czechia":                           "CZ",
	"Cote d'Ivoire":                     "CI",
	"Denmark":                           "DK",
	"Djibouti":                          "DJ",
	"Dominica":                          "DM",
	"Dominican Republic":                "DO",
	"Ecuador":                           "EC",
	"Egypt":                             "EG",
	"El Salvador":                       "SV",
	"Equatorial Guinea":                 "GQ",
	"Eritrea":                           "ER",
	"Estonia":                           "EE",
	"Eswatini":                          "SZ",
	"Ethiopia":                          "ET",
	"Falkland Islands (Malvinas)":       "FK",
	"Faroe Islands":                     "FO",
	"Fiji":                              "FJ",
	"Finland":                           "FI",
	"France":                            "FR",
	"French Guiana":                     "GF",
	"French Polynesia":                  "PF",
	"French Southern Territories":       "TF",
	"Gabon":                             "GA",
	"Gambia":                            "GM",
	"Georgia":                           "GE",
	"Germany":                           "DE",
	"Ghana":                             "GH",
	"Gibraltar":                         "GI",
	"Greece":                            "GR",
	"Greenland":                         "GL",
	"Grenada":                           "GD",
	"Guadeloupe":                        "GP",
	"Guam":                              "GU",
	"Guatemala":                         "GT",
	"Guernsey":                          "GG",
	"Guinea":                            "GN",
	"Guinea-Bissau":                     "GW",
	"Guyana":                            "GY",
	"Haiti":                             "HT",
	"Heard Island and McDonald Islands": "HM",
	"Vatican City":                      "VA",
	"Holy See":                          "VA",
	"Honduras":                          "HN",
	"Hong Kong":                         "HK",
	"Hungary":                           "HU",
	"Iceland":                           "IS",
	"India":                             "IN",
	"Indonesia":                         "ID",
	"Iran":                              "IR",
	"Iraq":                              "IQ",
	"Ireland":                           "IE",
	"Isle of Man":                       "IM",
	"Israel":                            "IL",
	"Italy":                             "IT",
	"Jamaica":                           "JM",
	"Japan":                             "JP",
	"Jersey":                            "JE",
	"Jordan":                            "JO",
	"Kazakhstan":                        "KZ",
	"Kenya":                             "KE",
	"Kiribati":                          "KI",
	"North Korea":                       "KP",
	"South Korea":                       "KR",
	"Kosovo":                            "XK",
	"Kuwait":                            "KW",
	"Kyrgyzstan":                        "KG",
	"Lao People's Democratic Republic":  "LA",
	"Laos":                              "LA",
	"Latvia":                            "LV",
	"Lebanon":                           "LB",
	"Lesotho":                           "LS",
	"Liberia":                           "LR",
	"Libya":                             "LY",
	"Liechtenstein":                     "LI",
	"Lithuania":                         "LT",
	"Luxembourg":                        "LU",
	"Macao":                             "MO",
	"Northern Macedonia":                "MK",
	"North Macedonia":                   "MK",
	"Madagascar":                        "MG",
	"Malawi":                            "MW",
	"Malaysia":                          "MY",
	"Maldives":                          "MV",
	"Mali":                              "ML",
	"Malta":                             "MT",
	"Marshall Islands":                  "MH",
	"Martinique":                        "MQ",
	"Mauritania":                        "MR",
	"Mauritius":                         "MU",
	"Mayotte":                           "YT",
	"Mexico":                            "MX",
	"Micronesia, Federated States of":   "FM",
	"Moldova":                           "MD",
	"Monaco":                            "MC",
	"Mongolia":                          "MN",
	"Montenegro":                        "ME",
	"Montserrat":                        "MS",
	"Morocco":                           "MA",
	"Mozambique":                        "MZ",
	"Myanmar":                           "MM",
	"Burma":                             "MM",
	"Namibia":                           "NA",
	"Nauru":                             "NR",
	"Nepal":                             "NP",
	"Netherlands":                       "NL",
	"New Caledonia":                     "NC",
	"New Zealand":                       "NZ",
	"Nicaragua":                         "NI",
	"Niger":                             "NE",
	"Nigeria":                           "NG",
	"Niue":                              "NU",
	"Norfolk Island":                    "NF",
	"Northern Mariana Islands":          "MP",
	"Norway":                            "NO",
	"Oman":                              "OM",
	"Pakistan":                          "PK",
	"Palau":                             "PW",
	"Palestine, State of":               "PS",
	"Panama":                            "PA",
	"Papua New Guinea":                  "PG",
	"Paraguay":                          "PY",
	"Peru":                              "PE",
	"Philippines":                       "PH",
	"Pitcairn":                          "PN",
	"Poland":                            "PL",
	"Portugal":                          "PT",
	"Puerto Rico":                       "PR",
	"Qatar":                             "QA",
	"Romania":                           "RO",
	"Russia":                            "RU",
	"Rwanda":                            "RW",
	"Reunion":                           "RE",
	"Saint Barthelemy":                  "BL",
	"Saint Helena":                      "SH",
	"Saint Kitts and Nevis":             "KN",
	"Saint Lucia":                       "LC",
	"Saint Martin (French part)":        "MF",
	"Saint Pierre and Miquelon":         "PM",
	"Saint Vincent and the Grenadines":  "VC",
	"Samoa":                             "WS",
	"San Marino":                        "SM",
	"Sao Tome and Principe":             "ST",
	"Saudi Arabia":                      "SA",
	"Senegal":                           "SN",
	"Serbia":                            "RS",
	"Seychelles":                        "SC",
	"Sierra Leone":                      "SL",
	"Singapore":                         "SG",
	"Sint Maarten (Dutch part)":         "SX",
	"Slovakia":                          "SK",
	"Slovenia":                          "SI",
	"Solomon Islands":                   "SB",
	"Somalia":                           "SO",
	"South Africa":                      "ZA",
	"South Georgia and the South Sandwich Islands": "GS",
	"South Sudan":                          "SS",
	"Spain":                                "ES",
	"Sri Lanka":                            "LK",
	"Sudan":                                "SD",
	"Suriname":                             "SR",
	"Svalbard and Jan Mayen":               "SJ",
	"Sweden":                               "SE",
	"Switzerland":                          "CH",
	"Syria":                                "SY",
	"Taiwan":                               "TW",
	"Tajikistan":                           "TJ",
	"Tanzania":                             "TZ",
	"Thailand":                             "TH",
	"Timor-Leste":                          "TL",
	"Togo":                                 "TG",
	"Tokelau":                              "TK",
	"Tonga":                                "TO",
	"Trinidad and Tobago":                  "TT",
	"Tunisia":                              "TN",
	"Turkey":                               "TR",
	"Turkmenistan":                         "TM",
	"Turks and Caicos Islands":             "TC",
	"Tuvalu":                               "TV",
	"Uganda":                               "UG",
	"Ukraine":                              "UA",
	"United Arab Emirates":                 "AE",
	"United Kingdom":                       "GB",
	"United States":                        "US",
	"United States Minor Outlying Islands": "UM",
	"Uruguay":                              "UY",
	"Uzbekistan":                           "UZ",
	"Vanuatu":                              "VU",
	"Venezuela":                            "VE",
	"Vietnam":                              "VN",
	"British Virgin Islands":               "VG",
	"US Virgin Islands":                    "VI",
	"Wallis and Futuna":                    "WF",
	"Western Sahara":                       "EH",
	"Yemen":                                "YE",
	"Zambia":                               "ZM",
	"Zimbabwe":                             "ZW",
}

type CombinedData struct {
	Id        string
	Country   string
	Date      string
	Infected  int
	Deaths    int
	Recovered int
}

type SingleData struct {
	Country string
	Date    string
	Value   int
}

// ParseCSVData parses the CSV data into a map with the country name as the key and the combined
// confirmed cases, total deaths, and recovered cases for each state within the country
func ParseCSVData(input io.Reader) (map[string]*SingleData, error) {
	data := make(map[string]*SingleData)
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
			data[country] = &SingleData{
				Country: country,
				Date:    mostRecentDate,
				Value:   mostRecentValue,
			}
		}
	}

	return data, err
}

// stringCleanup removes any special characters that exists within the string and
// if the country is broken up into two parts it corrects it (e.g Korea, South -> South Korea)
func stringCleanup(s string) string {
	if s == "US" {
		return "United States"
	}

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
func MergeResults(confirmedCases, deaths, recoveredCases map[string]*SingleData) []*CombinedData {
	var covidData []*CombinedData

	for _, confirmed := range confirmedCases {
		deadCases := deaths[confirmed.Country]
		recovered := recoveredCases[confirmed.Country]

		country := stringCleanup(confirmed.Country)
		id := countryIds[country]

		newData := &CombinedData{
			Id:        id,
			Country:   country,
			Date:      confirmed.Date,
			Infected:  confirmed.Value,
			Deaths:    deadCases.Value,
			Recovered: recovered.Value,
		}

		covidData = append(covidData, newData)
	}

	// Sort from most to least deaths
	sort.Slice(covidData[:], func(i, j int) bool {
		return covidData[i].Deaths > covidData[j].Deaths
	})

	return covidData
}

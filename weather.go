package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const URL = "https://api.openweathermap.org/data/2.5/weather"

//const URL = "https://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=3a17a33b2c242e1bc22f0303ff5c7ebc"

// API holds the http client and the url to call
type API struct {
	Client  *http.Client
	Key     string
	City    string
	Country string
}

// Wind is the current wind
type Wind struct {
	Speed   float64 `json:"speed"`
	Degrees int     `json:"deg"`
}

// Simple is the simple weather description
type Simple struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Temperature is current temperature
type Temperature struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
}

// Current holds the current weather (response from api)
type Current struct {
	Temperature Temperature `json:"main"`
	Simple      []Simple    `json:"weather"`
	Wind        Wind        `json:"wind"`
}

// CurrentWeather return current weather
func (api *API) CurrentWeather() (Current, error) {
	var cw Current
	link := fmt.Sprintf("%s?q=%s,%s&APPID=%s", URL, api.City, api.Country, api.Key)

	resp, err := api.Client.Get(link)
	if err != nil {
		log.Println(err)
		return cw, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&cw)
	if err != nil {
		log.Println(err)
		return cw, err
	}

	cw.Temperature.Temp = kelvinToCelsius(cw.Temperature.Temp)
	cw.Temperature.FeelsLike = kelvinToCelsius(cw.Temperature.FeelsLike)

	return cw, nil
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

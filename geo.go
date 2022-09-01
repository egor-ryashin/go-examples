package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Location struct {
	Longitude float64
	Latitude  float64
}

type Geo struct {
	Location Location
}

func geo() {

	url := "https://ip-geo-location.p.rapidapi.com/ip/check?format=json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("geo_key"))
	req.Header.Add("X-RapidAPI-Host", "ip-geo-location.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var geo Geo
	json.Unmarshal(body, &geo)

	url = fmt.Sprintf("https://everyearthquake.p.rapidapi.com/recentEarthquakes?interval=P1M&start=1&count=10&type=earthquake&latitude=%f&longitude=%f&radius=1000&units=miles", geo.Location.Latitude, geo.Location.Longitude)

	req, _ = http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("geo_key"))
	req.Header.Add("X-RapidAPI-Host", "everyearthquake.p.rapidapi.com")

	res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

}

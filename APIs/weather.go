package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string = "http://api.weatherapi.com/v1/current.json?key=d1f5fdb4dfb34c6aa9833229211008&q="

type Payload struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type Current struct {
	Temp      float32 `json:"temp_f"`
	Windspeed float32 `json:"wind_mph"`
}

func consume(s string) {
	response, err := http.Get(url + s)
	if err != nil {
		panic(err.Error())
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(data))
	var p Payload
	json.Unmarshal(data, &p)
	fmt.Println("\n", "You are skating in:", p.Location.Name+",", p.Location.Region, "\n", "Current temperature is:", p.Current.Temp, "degrees", "\n",
		"There is a wind speed of", p.Current.Windspeed, "Mph")
}

func Main() {
	answer := ""
	fmt.Println("Enter your zipcode for a quick weather report!")
	fmt.Scanf("%s", &answer)
	consume(answer)

}

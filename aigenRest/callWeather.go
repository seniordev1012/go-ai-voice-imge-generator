package aigenRest

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCurrentWeather() string {

	url := "https://api.openweathermap.org/data/2.5/weather?lat=-22.954531&lon=30.469860&appid=b82cdea422d224c48f04cb22e84bf279"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return "Error getting weather" + err.Error()
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "Error getting weather, with error" + err.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "Error getting weather, with error" + err.Error()
	}
	log.Println(string(body))
	return string(body)
}

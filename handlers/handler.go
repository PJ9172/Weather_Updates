package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func LandingPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func ShowWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	url := "https://api.openweathermap.org/data/2.5/weather?appid=" + os.Getenv("API_ID") + "&units=metric&q=" + city
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in Get method!!! : ", err)
		return
	}
	defer res.Body.Close()

	byteData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in Reading data!!! : ", err)
		return
	}

	temp, err := template.ParseFiles("templates/result.html")
	if err != nil {
		fmt.Println("Error to load template : ", err)
		return
	}
	temp.Execute(w, template.JS(byteData))
}

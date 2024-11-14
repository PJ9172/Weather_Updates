package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", landingPage)
	http.HandleFunc("/weather/", showWeather)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server starts on Port:3000!!!")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Errro to start Server!!!")
		return
	}
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func showWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	url := "https://api.openweathermap.org/data/2.5/weather?appid=7c801c838f8f1bacdcee99420a82e2e1&units=metric&q=" + city
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

	temp, err := template.ParseFiles("result.html")
	if err != nil {
		fmt.Println("Error to load template : ", err)
		return
	}
	temp.Execute(w, template.JS(byteData))
}

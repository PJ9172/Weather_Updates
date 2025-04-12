package main

import (
	"Weather_Tracking/config"
	"Weather_Tracking/handlers"
	"fmt"
	"net/http"
)

func main() {

	config.LoadEnv()
	http.HandleFunc("/", handlers.LandingPage)
	http.HandleFunc("/weather/", handlers.ShowWeather)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server starts on Port:3000!!!")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Errro to start Server!!!")
		return
	}
}

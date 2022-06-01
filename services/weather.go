package services

import (
	"go-autoreload/models"
	"html/template"
	"net/http"
)

func GetWeather() *models.Weather {
	water := Randomize()
	wind := Randomize()
	statusWater := GetStatusWater(water)
	statusWind := GetStatusWind(wind)

	weather := models.Weather{
		Water:       water,
		Wind:        wind,
		StatusWater: statusWater,
		StatusWind:  statusWind,
	}

	return &weather
}

func StartWeather(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html")

	tpl, err := template.ParseFiles("./html/template.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	weather := GetWeather()

	var status = map[string]interface{}{
		"water":        weather.Water,
		"wind":         weather.Wind,
		"status_water": weather.StatusWater,
		"status_wind":  weather.StatusWind,
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = tpl.Execute(rw, status)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	return

}

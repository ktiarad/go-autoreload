package models

type Weather struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

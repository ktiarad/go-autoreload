package services

func GetStatusWater(val int) string {
	if val < 5 {
		return "Aman"
	} else if val >= 6 && val <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func GetStatusWind(val int) string {
	if val < 6 {
		return "Aman"
	} else if val >= 7 && val <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetLatLon(city string) (float64, float64, error) {
	logger.Info(fmt.Sprintf("city: %s", city ))
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", city)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	var results []struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return 0, 0, err
	}
	if len(results) == 0 {
		return 0, 0, fmt.Errorf("no results found")
	}
	lat, err := strconv.ParseFloat(results[0].Lat, 64)
	if err != nil {
		return 0, 0, err
	}
	lon, err := strconv.ParseFloat(results[0].Lon, 64)
	if err != nil {
		return 0, 0, err
	}
	return lat, lon, nil
}

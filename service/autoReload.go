package service

import (
	"assigment3/structs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func AutoReload() {
	for {
		water := randomNumberWater()
		wind := randomNumberWind()

		num := structs.StatusWaterWind{}
		num.Status.Water = water
		num.Status.Wind = wind

		jsonData, err := json.Marshal(num)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = ioutil.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

func ReloadWeb(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	var status structs.StatusWaterWind

	err = json.Unmarshal(jsonData, &status)
	if err != nil {
		log.Fatal(err.Error())
	}

	water := status.Status.Water
	wind := status.Status.Wind

	var (
		statusWater string
		statusWind  string
	)

	statusWater = Water(water)
	statusWind = Wind(wind)

	data := map[string]interface{}{
		"statusWater": statusWater,
		"statusWind":  statusWind,
		"water":       water,
		"wind":        wind,
	}
	return (w,data)
}

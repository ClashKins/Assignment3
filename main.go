package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Indicator indicator `json:"status"`
}

type indicator struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func autohit() {
	for {
		min := 1
		max := 15
		rand.Seed(time.Now().UnixNano())
		numBrandWater := rand.Intn(max - min)
		numBrandWind := rand.Intn(max - min)

		indicator := Status{
			Indicator: indicator{
				Water: numBrandWater,
				Wind:  numBrandWind,
			},
		}
		jsonprint, err := json.MarshalIndent(indicator, "", " ")
		if err != nil {
			fmt.Println("json print error")
			return
		}
		log.Println(string(jsonprint))

	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("error making an avatar")
		return
	}

	var data Status
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Println("error unmarshal data")
			return
		}

	var waterStatus string
	var windStatus string
	water := data.Indicator.Water
	wind := data.Indicator.Wind

	if water > 8 {
		waterStatus = "bahaya"
	} else if water > 5 {
		waterStatus = "siaga"
	} else {
		waterStatus = "aman"
	}

	if wind > 15 {
		windStatus = "bahaya"
	} else if wind > 6 {
		windStatus = "siaga"
	} else {
		windStatus = "aman"
	}

	response := map[string]interface{}{
		"water":       water,
		"wind":        wind,
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
	}
	fmt.Println(response)

		err = ioutil.WriteFile("data.json", jsonprint, 0644)
		if err != nil {
			log.Fatalln("error auto reload data.json file", err)
		}

		time.Sleep(15 * time.Second)
	}
}

func main() {
	go autohit()

	// http.HandleFunc("/", GetAvatar)

	fmt.Println("Your server will be serve at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// func GetAvatar(w http.ResponseWriter, r *http.Request) {
// 	jsonData, err := ioutil.ReadFile("data.json")
// 	if err != nil {
// 		fmt.Println(w, "error making an avatar")
// 		return
// 	}

	// var data Status
	// err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Fprintln(w, "error unmarshal data")
	// 	return
	// }

// 	var waterStatus string
// 	var windStatus string
// 	water := data.Indicator.Water
// 	wind := data.Indicator.Wind

// 	if water > 8 {
// 		waterStatus = "bahaya"
// 	} else if water > 5 {
// 		waterStatus = "siaga"
// 	} else {
// 		waterStatus = "aman"
// 	}

// 	if wind > 15 {
// 		windStatus = "bahaya"
// 	} else if wind > 6 {
// 		windStatus = "siaga"
// 	} else {
// 		windStatus = "aman"
// 	}

// 	response := map[string]interface{}{
// 		"water":       water,
// 		"wind":        wind,
// 		"waterStatus": waterStatus,
// 		"windStatus":  windStatus,
// 	}
// 	fmt.Println(response)

// }

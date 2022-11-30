package main

import (
	"encoding/json"
	"fmt"
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
		max := 100
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
		
	switch numBrandWater > 0 {
	case numBrandWater < 5:
		result := "aman"
		fmt.Printf("status water : %s\n\n", result)
	case numBrandWater > 5 && numBrandWater < 8:
		result := "siaga"
		fmt.Printf("status water : %s\n\n", result)
	case numBrandWater > 8:
		result := "bahaya"
		fmt.Printf("status water : %s\n\n", result)
	default:
		result := "water measurable error"
		fmt.Printf("%s\n\n", result)
	}
	switch numBrandWind > 0 {
	case numBrandWind < 6 :
		result := "aman"
		fmt.Printf("wind status :%s\n\n", result)
	case numBrandWind > 6 && numBrandWind < 16:
		result := "siaga"
		fmt.Printf("wind status :%s\n\n", result)
	case numBrandWind > 15:
		result := "bahaya"
		fmt.Printf("wind status :%s\n\n", result)
	default:
		result := "wind measurable error"
		fmt.Printf("%s\n\n", result)
	}
		time.Sleep(15 * time.Second)
	}
	
func main() {
	go autohit()

	// http.HandleFunc("/", GetAvatar)

	fmt.Println("Your server will be serve at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
	

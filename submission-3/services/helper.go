package services

import (
	"encoding/json"
	"math/rand"
	"os"
)

func GenerateRandomInt() int {
	return rand.Intn(100)
}

func ParseResponse(data Data) []Response {
	waterStatus := "Aman"
	if data.Status.Water > 5 && data.Status.Water < 9 {
		waterStatus = "Siaga"
	} else if data.Status.Water > 8 {
		waterStatus = "Bahaya"
	}

	windStatus := "Aman"
	if data.Status.Wind > 6 && data.Status.Wind < 16 {
		windStatus = "Siaga"
	} else if data.Status.Wind > 15 {
		windStatus = "Bahaya"
	}
	return []Response{
		{
			Name:   "Water",
			Status: waterStatus,
			Value:  data.Status.Water,
		},
		{
			Name:   "Wind",
			Status: windStatus,
			Value:  data.Status.Wind,
		},
	}

}
func ReadJSONFile() Data {
	var data Data
	var text = make([]byte, 1024)

	WriteJSONFile()
	text, err := os.ReadFile(PATH_JSON)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(text, &data)

	return data

}

func WriteJSONFile() {
	randomData := RandomData()
	jsonData, err := json.MarshalIndent(randomData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(PATH_JSON, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func RandomData() Data {
	return Data{
		Status: Status{
			Water: GenerateRandomInt(),
			Wind:  GenerateRandomInt(),
		},
	}
}

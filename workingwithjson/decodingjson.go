package main

import (
	"encoding/json"
	"fmt"
	"strings"
	//"io"
)

func main() {
	reader := strings.NewReader(`{"Kayak":279, "LifeJacket":49.95}`)

	m := map[string]float64{}

	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println("Error: %v", err.Error())
	} else {
		fmt.Println("Map: %T, %v", m, m)
		for k, v := range m {
			fmt.Println("Key: %v, Value: %v", k, v)
		}
	}
}

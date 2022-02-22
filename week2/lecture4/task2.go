package main

import (
	"fmt"
	"math/rand"
	"time"
)



func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	// cityChoices := []string{"Berlin", "Tokyo"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}
func groupSlices(keySlice []string, valueSlice []int) map[string][]int {
	result := make(map[string][]int)

	for i, city := range keySlice {
		result[city] = append(result[city], valueSlice[i])

		
	}

	return result
}

func main() {
	cities, prices := citiesAndPrices()
	// cities := []string{"Moscow", "NewYork", "Moscow"}
	// prices := []int{4,2, 2}
	fmt.Println(groupSlices(cities, prices))
	
}
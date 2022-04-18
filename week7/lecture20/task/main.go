package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Drink struct {
	Name string `json:"strDrink"`
	Instructions string `json:"strInstructions"`
}

type DrinksResponsePayload struct {
	Drinks []Drink
}


type CocktailBartender struct {
	Data Drink
	url string

}

func (b *CocktailBartender) DoWork() {
	req, err := http.NewRequest("GET", b.url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	payload := DrinksResponsePayload{}
	json.NewDecoder(resp.Body).Decode(&payload)
	// fmt.Println(payload.Drinks[0].Name)
	// fmt.Println(payload.Drinks[0].Instructions)
	var currentDrink Drink
	if len(payload.Drinks) > 0 {
		currentDrink = payload.Drinks[0]

	}

	fmt.Println("Name: " + currentDrink.Name)

	var instructions = strings.Split(currentDrink.Instructions, `.`)
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}


	b.Data = currentDrink
	defer resp.Body.Close()

}

func (b *CocktailBartender) Start() {
	for {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("Choose your drink")
		var drink string
		fmt.Scan(&drink)
		if (drink == "nothing"){
			break
		}
		b.url = "https://www.thecocktaildb.com/api/json/v1/1/search.php?s=" + drink 
		b.DoWork()
	}
	fmt.Println("End of program")
}

func main() {

	b := &CocktailBartender{}
	b.Start()
	// fmt.Println(b.Data)
	
}
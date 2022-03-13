package main

import (
	"fmt"
	"log"

	"github.com/MichaelGenchev/scalefocusAcademy2022/cardgame"

	"github.com/MichaelGenchev/scalefocusAcademy2022/carddraw"
)

func main() {


	deck := cardgame.New()
	drawed, err := carddraw.DrawAllCards(&deck)
	if err != nil {
		log.Fatal(err)
		
	}else {
		fmt.Println(drawed)
	}
}
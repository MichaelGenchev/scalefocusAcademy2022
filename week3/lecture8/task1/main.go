package main

import (
	"fmt"

	"github.com/MichaelGenchev/scalefocusAcademy2022/cardgame"

	"github.com/MichaelGenchev/scalefocusAcademy2022/carddraw"

)

func main() {


	deck := cardgame.New()
	drawed := carddraw.DrawAllCards(&deck)

	fmt.Println(len(drawed))
}




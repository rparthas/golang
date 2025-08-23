package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (g gasEngine) rangeInMiles() uint16 {
	return uint16(g.mpg) * uint16(g.gallons)
}

func main() {
	var engine1 = gasEngine{mpg: 25, gallons: 15}
	fmt.Println(engine1.mpg, engine1.gallons, engine1.rangeInMiles())
}

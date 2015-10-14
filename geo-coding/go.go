package main

import (
	"log"

	"github.com/kellydunn/golang-geo"
)

func main() {
	p0 := geo.NewPoint(0, 0)
	p1 := geo.NewPoint(0, 1)
	traveledDistance := p0.GreatCircleDistance(p1)
	log.Println(traveledDistance)

	p := geo.NewPoint(49.014, 8.4043)
	geocoder := new(geo.GoogleGeocoder)
	res, err := geocoder.ReverseGeocode(p)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(res))

	// This gives an error:
	pnull := geo.NewPoint(0.0, 0.0)
	res, err = geocoder.ReverseGeocode(pnull)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(res))
}

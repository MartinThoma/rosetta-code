package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kr/pretty"
)

type search struct {
	GasStations []GasStation `xml:"gasStations>gasStation"`
	Forecast    Forecast     `xml:"forecast"`
}

type Forecast struct {
	trend     string `xml:"trend,attr"`
	cdataText string `xml:"cdataText,chardata"`
}

type GasStation struct {
	//Fuels    []Fuel `xml:"fuels>"`
	XMLName  xml.Name `xml:"gasStation"`
	Ident    string   `xml:"ident,attr"`
	Brand    string   `xml:"brand,attr"`
	Name     string   `xml:"name,attr"`
	Street   string   `xml:"street,attr"`
	Zip      string   `xml:"zip,attr"`
	City     string   `xml:"city,attr"`
	Lat      float64  `xml:"lat,attr"`
	Lng      float64  `xml:"lng,attr"`
	Highway  string   `xml:"highway,attr"`
	Services string   `xml:"services,attr"`
	Distance float64  `xml:"distance,attr"`
	IsOpen   bool     `xml:"is_open,attr"`
	Fuels    []Fuel   `xml:"fuels>fuel"`
}

type Fuel struct {
	XMLName   xml.Name `xml:"fuel"`
	Kind      string   `xml:"kind,attr"`
	Price     float64  `xml:"price,attr"`
	Timestamp int64    `xml:"timestamp,attr"`
}

func main() {
	xmlFile, err := os.Open("searchresult.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	var c search
	xml.Unmarshal(XMLdata, &c)

	fmt.Printf("%# v", pretty.Formatter(c.GasStations[0]))
}

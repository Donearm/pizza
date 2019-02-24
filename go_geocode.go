package main

////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2019, Gianluca Fiore
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
////////////////////////////////////////////////////////////////////////////////


import (
	"fmt"
	"encoding/json"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/mziia/geogoth"
)

type pizzeria struct {
	address string
	city string
	name string
	website string
}

func newCollection(c *geogoth.Features, lat float64, lng float64, id string, p *pizzeria) {
	point := geogoth.NewPoint([]float64{lat, lng})
	feature := geogoth.NewFeature()

	feature.SetID(id)
	feature.SetProperty("city", p.city)
	feature.SetProperty("pizzeria", p.name)
	feature.SetProperty("address", p.address)
	feature.SetProperty("website", p.website)
	feature.SetGeometry(point)
	c.AddFeature(feature)

	m, _ := json.Marshal(c)
	fmt.Println(string(m))
}


func main() {
	// Initialize OpenStreetMap geocoder
	g := openstreetmap.Geocoder()
	// Initialize test pizzeria
	n := pizzeria{"623 E. Adams St.", "Phoenix", "Pizzeria Bianco", "http://www.pizzeriabianco.com"}
	testAddress := n.address + ", " + n.city
	location, _ := g.Geocode(testAddress)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", testAddress, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := g.ReverseGeocode(location.Lat, location.Lng)
	if address != nil {
		fmt.Printf("Address of (%.6F, %.6f) is %s\n", location.Lat, location.Lng, address.FormattedAddress)
		fmt.Printf("Detailed address: %#v\n", address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Print("\n")


	featureC := geogoth.NewFeatureCollection()
	newCollection(featureC, 40.55555, 11.4343, "00001", &n)
}


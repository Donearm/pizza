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

	"github.com/codingsince1985/geo-golang/openstreetmap"
)

const (
	ADDR = "Rydlowka 31/46, 30363, Krakow, Poland"
)

func main() {
	g := openstreetmap.Geocoder()
	location, _ := g.Geocode(ADDR)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", ADDR, location.Lat, location.Lng)
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
}


package main

////////////////////////////////////////////////////////////////////////////////
//
// Copyright (c) 2019-2020, Gianluca Fiore
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"encoding/json"
	"os"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/mziia/geogoth"

	"./pizzerias" // load local PizzaList and Pizzeria struct
)

const GEOJSON_FILE = "pizzerias.geojson"

// Add a new feature to a collection
func newCollection(c *geogoth.Features, lat float64, lng float64, id string, p *pizzerias.Pizzeria) {
	point := geogoth.NewPoint([]float64{lng, lat})
	feature := geogoth.NewFeature()

	feature.SetID(id)
	feature.SetProperty("city", p.City)
	feature.SetProperty("pizzeria", p.Name)
	feature.SetProperty("address", p.Address)
	feature.SetProperty("website", p.Website)
	feature.SetGeometry(point)
	c.AddFeature(feature)
}

// Save a GeoJSON byte array into a file
func saveGeoJsonToFile(jsondata []byte) {
	f, err := os.Create(GEOJSON_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err2 := f.Write(jsondata)
	if err2 != nil {
		panic(err2)
	}
}

func main() {
	// Initialize GeoJSON collection
	featureC := geogoth.NewFeatureCollection()
	// Initialize OpenStreetMap geocoder
	g := openstreetmap.Geocoder()
	// Initialize each pizzeria
	for i, n := range pizzerias.PizzaList {
		testAddress := n.Address + ", " + n.City
		location, _ := g.Geocode(testAddress)
		if location == nil {
			fmt.Println("got <nil> location")
			fmt.Println("Pizzeria that couldn't be found was ", testAddress)
			continue
		}
		address, _ := g.ReverseGeocode(location.Lat, location.Lng)
		if address != nil {
			fmt.Printf("Address of (%.6F, %.6f) is %s\n", location.Lat, location.Lng, address.FormattedAddress)
		} else {
			continue
		}

		// Save the just found point on map into a new feature of the GeoJSON 
		// collection
		newCollection(featureC, location.Lat, location.Lng, string(i), &n)
	}
	// Have the collection made into GeoJSON
	m, _ := json.Marshal(featureC)
	// Save the final collection to a file
	saveGeoJsonToFile(m)
}

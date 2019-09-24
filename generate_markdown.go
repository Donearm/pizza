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
	"os"

	"./pizzerias"
)

const output = "README.md"
const pizzaTitle = "# Pizza"
const pizzaGif = `<p align="center">
	<img src="pizza.gif" type="image/gif" alt="where's the pizza?">
</p>
`
const pizzaDesc = `Where is the best pizza in the world? 

This is a community project that couldn't be born out of a single person and 
thus collects the contributions of many collaborators. Originally it was at <a 
href="https://github.com/stevekinney/pizza" alt="original 
project">stevekinney/pizza</a>, to which go my thanks for the idea and first 
realization.

It goes without saying that many cities and countries are missing and we need 
your contributions to signal the best pizzerias in them. Only true pizza 
though, no deep-dishes or focaccias. Pull requests are more than welcome, vital 
even.`

const pizzaContributing = `To actually know how to contribute, check the [Contribute file](CONTRIBUTING)`

// Helper function to write a string to a file
func writeToFile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error in creating the readme")
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.Write([]byte(data)); err != nil {
		fmt.Println("Error in writing a string to file: ", data)
		os.Exit(1)
	}

	file.Sync()
}


func main() {

	// Save to readme the description of the project
	writeToFile(output, pizzaTitle)
	writeToFile(output, "\n\n")
	writeToFile(output, pizzaGif)
	writeToFile(output, "\n")
	writeToFile(output, pizzaDesc)
	writeToFile(output, "\n\n")
	writeToFile(output, pizzaContributing)
	writeToFile(output, "\n\n")

	// output to readme the full list of pizzerias
	for _, p := range pizzerias.PizzaList {
		pizzaString := "* [" + p.Name + "](" + p.Website + ")" + " - " + p.Address + ", **" + p.City + "**, " + p.Country
		writeToFile(output, pizzaString)
		writeToFile(output, "\n")
	}
}


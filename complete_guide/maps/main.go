package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	//var cars map[string]string
	years := make(map[string]int)

	years["Toyota"] = 2020
	years["Honda"] = 2019

	delete(colors, "red")
	delete(years, "Honda")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

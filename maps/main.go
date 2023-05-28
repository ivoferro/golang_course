package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	colors["white"] = "#FFFFFF"
	colors["black"] = "#000000"

	PrintColors(colors)
}

func PrintColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}
}

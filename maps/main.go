package main

import "fmt"

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, val)
	}
}

func main() {

	colors := map[string]string{
		"red":   "#redhex",
		"green": "#greenhex",
		"white": "#whitehex",
		"pink":  "#pinkHex",
	}

	printMap(colors)

}

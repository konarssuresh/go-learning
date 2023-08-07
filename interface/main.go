package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	sb := spanishBot{}
	eb := englishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

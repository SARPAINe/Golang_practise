package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// we can omit eb if we are not using it
func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hola"
}

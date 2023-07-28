package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
	//we can also write just contactInfo
}

func main() {
	//  alex:=person{"Alex","Benjamin"}
	shaharin := person{firstName: "Shaharin", lastName: "Ahmed"}
	var reyad person
	reyad.firstName = "Shakil"
	reyad.lastName = "Ahmed"
	//  fmt.Println(alex.firstName,shaharin.lastName)
	shaharin.lastName = "Kazi"
	fmt.Println(shaharin)
	fmt.Printf("%+v\n", reyad)

	jim := person{firstName: "Jim", lastName: "Collins", contactInfo: contactInfo{email: "jim@gmail.com", zipCode: 940}}
	jim.print()
	jim.updateName("Shawn")
	jim.print()

	message := []string{"hi", "there"}
	fmt.Println(message)
	updateMessage(message)
	fmt.Println(message)

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateMessage(m []string) {
	m[0] = "bye"
}

// int, float, string, bool and structs are value types
// use pointers to change these in a function
// slices, maps, channels, pointers, functions inside functions
// these are reference types no need to use pointer.

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func main() {
	suresh := person{
		firstName: "suresh",
		lastName:  "konar",
		contactInfo: contactInfo{
			email:   "test@test.com",
			zipCode: 123,
		},
	}

	suresh.print()
	suresh.updateName("ashwini")
	suresh.print()
}

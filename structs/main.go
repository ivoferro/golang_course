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

func main() {
	ivo := person{
		firstName: "Ivo",
		lastName:  "Ferro",
		contactInfo: contactInfo{
			email:   "lez.go@leet.com",
			zipCode: 4470,
		},
	}
	//(&ivo).updateName("GrowLeR")
	ivo.updateName("GrowLeR")
	ivo.print()

	var ola person
	ola.firstName = "Ola"
	ola.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

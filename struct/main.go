package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// surya := person{ "surya",  "prasanth"}
	// surya := person{firstName: "surya", lastName: "prasanth"}
	// var surya person
	// surya.firstName = "surya"
	// surya.lastName = "prasanth"
	surya := person{firstName: "surya", lastName: "prasanth", contactInfo: contactInfo{email: "surya@gmail.com", zip: 1}}
	// surya.print()
	// suryaPointer := &surya
	surya.updateName("K Surya")
	surya.print()

}
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) print() {
	fmt.Println(*p)
}

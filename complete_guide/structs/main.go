package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

/*
type person struct {
	firstName string
	lastName  string
	contact  contactInfo
}
*/

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	//jack := person{"Jack", "Johnson"}
	var allison person

	fmt.Println(alex)
	fmt.Println(allison)

	fmt.Printf("%+v\n", allison)


	// Assigning values to the fields of the allison variable
	allison.firstName = "Allison"
	allison.lastName = "Anderson"

	fmt .Println(allison)
	fmt.Printf("%+v\n", alex)
	/*
	jim := person{
		firstName: "Jim",
		lastName:  "Johnson",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
    */
	jim := person{
		firstName: "Jim",
		lastName:  "Johnson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}


	fmt.Println(jim)
	fmt.Printf("%+v\n", jim)

	jim.print()
	jim.updateName("Jimmy")
	fmt.Printf("%+v\n", jim)

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}	


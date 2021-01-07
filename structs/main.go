package main

type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactinfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	jim.updateName("jimmy")
	//jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

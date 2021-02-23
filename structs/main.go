package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}


type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	//alex := person{
	//	firstName: "Udit",
	//	lastName:  "Gurani",
	//}
	//fmt.Println(alex)
	//
	//var alexOne person
	//alexOne.firstName = "UditOne"
	//alexOne.lastName = "SecondValue"
	//fmt.Println(alexOne)

	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact:   contactInfo{
			email:   "jim.andeson@email.com",
			zipCode: 0,
		},
	}
	//jim.print()
	//jimPointer := &jim
	jim.updateName("JimmyOne")
	fmt.Println(jim)
}

func (pointerTOPerson *person) updateName(newFirstName string)  {
	pointerTOPerson.firstName = newFirstName
}

//func (p person) print() {
//	fmt.Println("%+v",p)
//}

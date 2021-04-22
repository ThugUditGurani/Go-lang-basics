package main

import "fmt"

type user struct {
	name string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User EMail To %s<%s>\n",u.name,u.email)
}

func (u *user) changeEmail(s string) {
	u.email = s
}


func main() {
	//Values of Type user can be used to call methods
	//declared with a value receiver.
	bill := user{"Bill","bill@email.com"}
	bill.notify()

	//Pointers to type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"lisa", "lisa@email.com"}
	lisa.notify()

	//Values of type user can also bee used to call methods
	//declared with a pointer receiver
	bill.changeEmail("bill@nerdomain.com")
	bill.notify()

}

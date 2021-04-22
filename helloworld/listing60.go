package main

import "fmt"

// notifier is an interface that defined notification
// type behaviour
type notifier interface {
	notify()
}

// user defines a user in the program
type user23 struct {
	name string
	email string
}

// notify implements a method that can be called via
// a value of type user
func (u *user23) notify() {
	fmt.Printf("Sending user email to %s<%s> \n",u.name,u.email)
}


// Admin represents an admin user with privileges
type admin struct {
	user23
	level string
}

// notify implements a method that can be called via
// a value of type admin.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s> \n",a.name,a.email)
}


func main() {
	//Create an admin user.
	ad := admin{
		user23 : user23{
			name : "john smith",
			email: "john@email.com",
		},
		level: "super",
	}

	// Send the admin user a notification
	// The embedded inner type implementation of the
	// interface is NOT 'promoted' to the outer type
	sendNotification(&ad)

	// We can accesss the inner type method directly
	ad.user23.notify()

	// the inner type method is not promoted
	ad.notify()


}


func sendNotification(n notifier){
	n.notify()
}


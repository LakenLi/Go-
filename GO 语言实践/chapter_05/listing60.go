/*
* @Author: HoRan.li
* @Date:   2021-04-05 17:19:32
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-04-05 17:24:57
* 
*/
package main

import "fmt"

type notifier interface {
	notify()
}


type user struct {
	name   string
	email  string
}


func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>", u.name, u.email)
}


type admin struct {
	user
	level  string
}

func (a * admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>", a.name, a.email)
}


func main() {
	ad := admin {
		user: user {
			name: "afdf",
			email: "adfe@adsfda.com",
		},
		level: "super",
	}


	sendNotification(&ad)

	ad.user.notify()

	ad.notify()


}


func sendNotification(n notifier) {
	n.notify()
}
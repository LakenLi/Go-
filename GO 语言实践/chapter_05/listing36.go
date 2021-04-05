/*
* @Author: HoRan.li
* @Date:   2021-04-05 16:59:22
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-04-05 17:06:57
*/
package main

import "fmt"

type notifier interface {
	notify()
}


type user struct {
	name  string
	email string
}


func (u *user) notify() {
	fmt.Println("Sending user email to %s %s", u.name, u.email)
}


func main() {
	u := user{"Bill", "bill@email.com"}
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
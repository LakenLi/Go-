/*
* @Author: HoRan.li
* @Date:   2021-04-05 17:07:18
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-04-05 17:13:20
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

// notify 使用指针接收者实现了notifer接口
func (u *user)  notify() {
	fmt.Println("afd d s %s   %s", u.name, u.email)
}

type admin struct {
	name string
	email string
}

func (a *admin) notify() {
	fmt.Println("dfads %s %s", a.name, a.email)
}


func main() {
	bill := user{"bill", "bill@new.com"}
	sendNotification(&bill)

	lisa := admin{"lisa", "lisa@email.com"}
	sendNotification(&lisa)
}


func sendNotification(n notifier) {
	n.notify()
}

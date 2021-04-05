/*
* @Author: HoRan.li
* @Date:   2021-04-05 16:39:59
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-04-05 16:52:29
*/
package main

import "fmt"

// 定义一个用户类型
type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("bill@new.com")
	bill.notify()

	lisa.changeEmail(".com")
	lisa.notify()
}
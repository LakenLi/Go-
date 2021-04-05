/*
* @Author: HoRan.li
* @Date:   2021-03-25 22:36:31
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-03-25 22:38:17
*/
package main

import "fmt"

func main() {
	const LENGTH = 20
	const WIDTH = 10

	var area int

	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH

	fmt.Printf("面积为 : %d", area)
   	println()
   	println(a, b, c)
}
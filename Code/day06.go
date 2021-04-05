/*
* @Author: HoRan.li
* @Date:   2021-04-03 15:17:09
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-04-03 16:53:05
*/
package main

import "fmt"

func main() {
	var testMap map[string]string
	testMap = make(map[string]string)

	testMap["app"] = "gagda"
	testMap["app"] = "gagda233"
	testMap["app23"] = "gagdadfdf"

	defer fmt.Println("333")

	for ttt := range testMap {
		fmt.Println("caafd of", ttt, "is", testMap[ttt])
	}

	area, _ := rectProps(10, 20)

	defer fmt.Println("222")
	fmt.Println(area)


	rest := oper(100, 5, func(a, b int)int{
		if b == 0 {
			fmt.Println("除数不能为零")
			return 0
		}
		return a / b
	})

	fmt.Println(rest)
}



func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func oper(a, b int, fun func(int, int) int) int {
    fmt.Println(a, b, fun) //打印3个参数
    res := fun(a, b)
    return res
}
/*
* @Author: HoRan.li
* @Date:   2021-03-31 08:39:15
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-03-31 08:51:41
*/
package main
import "fmt"


func main(){
	source := []string{"apple", "orange", "plum", "banana", "grape"}

	slice := source[2:3:4]

	fmt.Println(slice)

	slice1 := source[2:3:3]
	fmt.Println(source)
	slice1 = append(slice1, "Kiwi")
	fmt.Println(source)
	fmt.Println(slice1)


	slice2 := [][]int{{10}, {100, 200}}
	slice2[0] = append(slice2[0], 20)
	fmt.Println(slice2)
}
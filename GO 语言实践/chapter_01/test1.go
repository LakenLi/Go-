/*
* @Author: HoRan.li
* @Date:   2021-03-21 22:55:55
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-03-21 23:29:53
*/
package main

import "fmt"


type person struct {  
    name string
}

func main() {
    // sn1 := struct {
    //     age  int
    //     name string
    // }{age: 11, name: "qq"}
    // sn2 := struct {
    //     age  int
    //     name string
    // }{age: 11, name: "qq"}

    // if sn1 == sn2 {
    //     fmt.Println("sn1 == sn2")
    // }

    // sm1 := struct {
    //     age int
    //     m   map[string]string
    // }{age: 11, m: map[string]string{"a": "1"}}
    // sm2 := struct {
    //     age int
    //     m   map[string]string
    // }{age: 11, m: map[string]string{"a": "1"}}

    // if sm1 == sm2 {
    //     fmt.Println("sm1 == sm2")
    // }
    

    // a := 5
    // b := 8.1
    // fmt.Println(a + b)
    // 
    

    // a := [2]int{5, 6}
    // b := [3]int{5, 6}
    // if a == b {
    //     fmt.Println("equal")
    // } else {
    //     fmt.Println("not equal")
    // }
    // 
    // 
    

    var m map[person]int
    p := person{"mike"}
    fmt.Println(p)
    fmt.Println(m[p])



}
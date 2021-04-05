/*
* @Author: HoRan.li
* @Date:   2021-03-25 22:47:09
* @Last Modified by:   HoRan.li
* @Last Modified time: 2021-03-29 21:40:11
*/
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)

    var A, B = 60, 13
    //var C = A & B
    fmt.Println(A &^ B)

    bar()
    baz()

    num := 75
    switch { // expression is omitted
	    case num >= 0 && num <= 50:
	        fmt.Println("num is greater than 0 and less than 50")
	    case num >= 51 && num <= 100:
	        fmt.Println("num is greater than 51 and less than 100")
	    case num >= 101:
	        fmt.Println("num is greater than 100")
    }


    /* 定义局部变量 */
   var aa int = 10

   /* 循环 */
   LOOP: for aa < 20 {
      if aa == 15 {
         /* 跳过迭代 */
         aa = aa + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", aa)
      aa++     
   }


    var a1 [4] float32 // 等价于：var arr2 = [4]float32{}
  	fmt.Println(a1) // [0 0 0 0]
  	var b1 = [5] string{"ruby", "王二狗", "rose"}
  	fmt.Println(b1) // [ruby 王二狗 rose  ]
  	var c1 = [5] int{'A', 'B', 'C', 'D', 'E'} // byte
  	fmt.Println(c1) // [65 66 67 68 69]
  	d1 := [...] int{1,2,3,4,5}// 根据元素的个数，设置数组的大小
  	fmt.Println(d1)//[1 2 3 4 5]
  	e1 := [5] int{4: 100} // [0 0 0 0 100]
  	fmt.Println(e1)
  	f1 := [...] int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
  	fmt.Println(f1)

}


func bar() {
    type ErrorCode int

    const (
        ERROR_SUCCESS ErrorCode = iota
        ERROR_FIRST
        ERROR_SECOND
        ERROR_THIRD
    )

    error_code := ERROR_SUCCESS
    fmt.Println("default: ", error_code) // default:  0

    error_code = ERROR_SECOND
    fmt.Println("Second: ", error_code) // Second:  2
}

func baz() {
    type Flags uint

    const (
        FlagUp Flags = 1 << iota // is up
        FlagBroadcast // 1 << 1 = 2^1 = 2, supports broadcast access capability
        FlagLoopback // 1 << 2 = 2 ^ 2 = 4, is a loopback interface
        FlagPointToPoint // 1 << 3 = 2 ^ 3 = 8, belongs to a point-to-point link
        FlagMulticast // 1 << 4 = 2 ^ 4 = 16, supports multicast access capability
    )

    fmt.Println("FlagUp: ", FlagUp)
    fmt.Println("FlagBroadcast: ", FlagBroadcast)
    fmt.Println("FlagLoopback: ", FlagLoopback)
    fmt.Println("FlagPointToPoint: ", FlagPointToPoint)
    fmt.Println("FlagMulticast: ", FlagMulticast)
}
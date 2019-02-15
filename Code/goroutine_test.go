package main

import (
	"fmt"
	"sync"
	"math"
)

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}

	fmt.Println(id, x)

}

func main() {
	wg := new (sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int){
			defer wg.Done()
			sum(id)
		}
	}

	wg.Wait()
}
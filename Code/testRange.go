package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sim:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "appale", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
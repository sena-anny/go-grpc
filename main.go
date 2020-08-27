package main

import (
	"fmt"
	"github.com/sena-anny/go-grpc/animals"

	_ "github.com/sena-anny/go-grpc/animals"
)

func main() {
	var p int
	a := &p
	p = 5
	fmt.Printf("typs = %T adress = %p value = %d \n",a, a, *a)
	*a = 10

	s := "PS"
	printString(s)

	n := Sum(
			[]int{1,2,3},
			func(i int) int {
				return i *2
			},
		)
	fmt.Println(n)

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.DogFeed())
}

func printString(s string) {
	fmt.Println(s)
}

type Callback func(i int) int
func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func test() string {
	return animals.ElephantFeed()
}

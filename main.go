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

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.DogFeed())
}

func printString(s string) {
	fmt.Println(s)
}

func test() string {
	return animals.ElephantFeed()
}

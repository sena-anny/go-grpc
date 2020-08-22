package main

import (
	"fmt"

	"github.com/sena-anny/go-grpc/animals"
)

func main() {
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.DogFeed())
}

func test() string {
	return animals.ElephantFeed()
}

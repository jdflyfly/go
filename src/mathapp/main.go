package main

import (
	"mymath"
	"fmt"
)

func main() {
	a :=100
	b :=200
	c :=a+b
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n%v", mymath.Sqrt(2),c)
}

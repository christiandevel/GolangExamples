package main

import (
	"examples/basico"
	"fmt"
)

func main() {
	fmt.Println("This repository is for learning golang.")
	fmt.Println("Basic Exercise *************************")
	basico.Helloword()
	basico.PrintNameAndAge("Christian", 24)
	basico.GetMultiplication(6)
	basico.FindMayorNumer([]int{31, 3, 10, 4, 5, 11, 1, 3, 4, 20, 8})
}

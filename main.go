package main

import (
	"fmt"

	augmat "github.com/cubexgani/quadratic-regression-calculator/utils"
)

func main() {
	co := [][]float32{
		{1, 1, 1},
		{2, -1, 1},
		{1, 0, 1},
	}
	val := []float32{6, 3, 4}

	noce, err := augmat.MakeAugMat(co, val)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(noce)
	noce.Reduce()
	fmt.Println("Reduced matrix")
	fmt.Println(noce)
	fmt.Println("Rank:", noce.Rank())
	solns := noce.GetSolutions()
	fmt.Println("Solution vector:", solns)
	fmt.Println("Hello warudo")
}

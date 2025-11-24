package main

import (
	"fmt"

	augmat "github.com/cubexgani/quadratic-regression-calculator/utils"
)

func main() {
	co := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	val := []int{10, 11, 12}

	noce := augmat.Make_aug_mat(co, val)
	fmt.Println(noce)
	fmt.Println("Hello warudo")
}

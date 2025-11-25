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
	solns, err := noce.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solution vector:", solns)
	}

	hco := [][]float32{
		{1, 3, -2},
		{2, -1, 4},
		{1, -11, 14},
	}
	hval := []float32{0, 0, 0}
	homo, err := augmat.MakeAugMat(hco, hval)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(homo)
	hsol, err := homo.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solution vector:", hsol)
	}

	nco := [][]float32{
		{1, 1, 1},
		{1, 1, 1},
		{2, -1, 1},
	}
	nval := []float32{3, 5, 1}
	nsol, err := augmat.MakeAugMat(nco, nval)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nsol)
	nslns, err := nsol.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solution vector:", nslns)
	}
	fmt.Println("Hello warudo")
}

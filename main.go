package main

import (
	"fmt"

	"github.com/cubexgani/regression-calculator/utils"
)

func main() {
	n, x, y, err := scanXY()

	if err != nil {
		return
	}

	xv, yv := utils.GetTable(n, x, y)
	// fmt.Println(xv)
	// fmt.Println(yv)

	co := [][]float32{
		{float32(xv.Num), xv.Sums[0], xv.Sums[1]},
		{xv.Sums[0], xv.Sums[1], xv.Sums[2]},
		{xv.Sums[1], xv.Sums[2], xv.Sums[3]},
	}
	val := []float32{yv.Sums[0], yv.Sums[1], yv.Sums[2]}

	noce, err := utils.MakeAugMat(co, val)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(noce)
	solns, err := noce.Solve()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Solution vector:", solns)
	}
	curve := utils.GetCurve(solns)
	fmt.Println("Curve:", curve)
	fmt.Println("Hello warudo")
}

func scanXY() (n int, x, y []float32, err error) {
	fmt.Print("Enter number of points: ")
	_, err = fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error while scanning n:", err)
		return
	}
	x = make([]float32, n)
	y = make([]float32, n)
	fmt.Println("Points:")
	for i := range n {
		_, err = fmt.Scan(&x[i], &y[i])
		if err != nil {
			fmt.Println("Error while scanning:", err)
			return
		}
	}
	// fmt.Printf("x: %v\ny:%v\n", x, y)
	return
}

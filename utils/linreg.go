package utils

import "fmt"

type LinReg struct {
	Num int
	xv  XVals
	yv  YVals
}

func GetLinTable(n int, x, y []float32) *LinReg {
	Powers_Xi = 2
	Powers_YXi = 2

	xp := make([][]float32, Powers_Xi)
	xs := make([]float32, Powers_Xi)
	yp := make([][]float32, Powers_YXi)
	ys := make([]float32, Powers_YXi)

	for i := range Powers_Xi {
		xp[i] = make([]float32, n)
	}
	for i := range Powers_YXi {
		yp[i] = make([]float32, n)
	}
	for i := range n {
		var prodXi float32 = 1
		for j := range Powers_Xi {
			prodXi *= x[i]
			xp[j][i] = prodXi
			xs[j] += prodXi
		}

		var prodYXi float32 = y[i]
		for j := range Powers_YXi {
			yp[j][i] = prodYXi
			ys[j] += prodYXi
			prodYXi *= x[i]
		}
	}
	xv := XVals{xp, xs}
	yv := YVals{yp, ys}
	return &LinReg{n, xv, yv}
}

func (lr *LinReg) Solve() ([]string, error) {

	co := [][]float32{
		{float32(lr.Num), lr.xv.Sums[0]},
		{lr.xv.Sums[0], lr.xv.Sums[1]},
	}
	val := []float32{lr.yv.Sums[0], lr.yv.Sums[1]}

	noce, err := MakeAugMat(co, val)
	if err != nil {
		return []string{}, err
	}
	solns, err := noce.Solve()
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	} else {
		fmt.Println("Solution vector:", solns)
	}
	return []string{lr.GetCurve(solns, 'y')}, nil
}

func (*LinReg) GetCurve(solns []float32, dependentVar rune) string {
	eqnStr := fmt.Sprintf("%c =", dependentVar)
	termCount := 0
	var independentVar rune
	if dependentVar == 'y' {
		independentVar = 'x'
	} else {
		independentVar = 'y'
	}

	if solns[0] != 0 {
		eqnStr = fmt.Sprintf("%s %.3f", eqnStr, solns[0])
		termCount++
	}
	if termCount == 0 {
		if solns[1] != 0 {
			eqnStr = fmt.Sprintf("%s %.3f%c", eqnStr, solns[1], independentVar)
			termCount++
		} else {
			eqnStr = fmt.Sprintf("%s %.3f", eqnStr, solns[1])
			termCount++
		}
	} else {
		if solns[1] > 0 {
			eqnStr = fmt.Sprintf("%s + %.3f%c", eqnStr, solns[1], independentVar)
		} else if solns[1] < 0 {
			eqnStr = fmt.Sprintf("%s - %.3f%c", eqnStr, -solns[1], independentVar)
		}
	}
	return eqnStr
}

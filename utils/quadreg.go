package utils

import "fmt"

type QuadReg struct {
	Num int
	XV  XVals
	YV  YVals
}

func GetQuadTable(n int, x, y []float32) *QuadReg {
	Powers_Xi = 4
	Powers_YXi = 3

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
	return &QuadReg{n, xv, yv}
}

func (qr *QuadReg) Solve() ([]string, error) {

	co := [][]float32{
		{float32(qr.Num), qr.XV.Sums[0], qr.XV.Sums[1]},
		{qr.XV.Sums[0], qr.XV.Sums[1], qr.XV.Sums[2]},
		{qr.XV.Sums[1], qr.XV.Sums[2], qr.XV.Sums[3]},
	}
	val := []float32{qr.YV.Sums[0], qr.YV.Sums[1], qr.YV.Sums[2]}

	noce, err := MakeAugMat(co, val)
	if err != nil {
		return []string{}, err
	}
	solns, err := noce.Solve()
	if err != nil {
		return []string{}, err
	}
	return []string{qr.GetCurve(solns, 'y')}, nil
}

func (*QuadReg) GetCurve(solns []float32, dependentVar rune) string {
	termCount := 0

	var independentVar rune
	// This program specifically assumes that if dependent variable is y, independent variable is x and vice versa
	if dependentVar == 'y' {
		independentVar = 'x'
	} else {
		independentVar = 'y'
	}

	eqnStr := fmt.Sprintf("%c =", dependentVar)

	if solns[0] != 0 {
		eqnStr = fmt.Sprintf("%s %.3f", eqnStr, solns[0])
		termCount++
	}
	for i := 1; i < len(solns); i++ {
		if termCount == 0 {
			if solns[i] != 0 {
				eqnStr = fmt.Sprintf("%s %.3f%c", eqnStr, solns[i], independentVar)
				termCount++
			} else {
				continue
			}
		} else if solns[i] < 0 {
			eqnStr = fmt.Sprintf("%s - %.3f%c", eqnStr, -solns[i], independentVar)
			termCount++
		} else if solns[i] > 0 {
			eqnStr = fmt.Sprintf("%s + %.3f%c", eqnStr, solns[i], independentVar)
			termCount++
		} else {
			continue
		}
		if i > 1 {
			eqnStr = fmt.Sprintf("%s^%d", eqnStr, i)
		}
	}
	if termCount == 0 {
		eqnStr = fmt.Sprintf("%s %.3f", eqnStr, solns[0])
	}
	return eqnStr
}

func (qr *QuadReg) GetData() (XVals, YVals) {
	return qr.XV, qr.YV
}

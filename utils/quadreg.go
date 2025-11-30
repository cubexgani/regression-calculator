package utils

import "fmt"

type QuadReg struct {
	Num int
	xv  XVals
	yv  YVals
}

func GetQuadTable(n int, x, y []float32) *QuadReg {
	xp := make([][]float32, 4)
	xs := make([]float32, 4)
	yp := make([][]float32, 3)
	ys := make([]float32, 3)
	for i := range 4 {
		xp[i] = make([]float32, n)
	}
	for i := range 3 {
		yp[i] = make([]float32, n)
	}
	for i := range n {
		xp[0][i] = x[i]
		xp[1][i] = x[i] * x[i]
		xp[2][i] = x[i] * x[i] * x[i]
		xp[3][i] = x[i] * x[i] * x[i] * x[i]

		xs[0] += xp[0][i]
		xs[1] += xp[1][i]
		xs[2] += xp[2][i]
		xs[3] += xp[3][i]

		yp[0][i] = y[i]
		yp[1][i] = x[i] * y[i]
		yp[2][i] = x[i] * x[i] * y[i]

		ys[0] += yp[0][i]
		ys[1] += yp[1][i]
		ys[2] += yp[2][i]

	}
	xv := XVals{xp, xs}
	yv := YVals{yp, ys}
	return &QuadReg{n, xv, yv}
}

func (qr *QuadReg) Solve() ([]string, error) {

	co := [][]float32{
		{float32(qr.Num), qr.xv.Sums[0], qr.xv.Sums[1]},
		{qr.xv.Sums[0], qr.xv.Sums[1], qr.xv.Sums[2]},
		{qr.xv.Sums[1], qr.xv.Sums[2], qr.xv.Sums[3]},
	}
	val := []float32{qr.yv.Sums[0], qr.yv.Sums[1], qr.yv.Sums[2]}

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
	return qr.GetCurve(solns), nil
}

//TODO: Eliminate cases like ... + -2x + -4.564x^2
func (*QuadReg) GetCurve(solns []float32) []string {
	return []string{fmt.Sprintf("y = %.3f + %.3fx + %.3fx^2", solns[0], solns[1], solns[2])}
}

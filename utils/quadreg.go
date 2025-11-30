package utils

import "fmt"

type QuadReg struct {
	Num int
	xv  XVals
	yv  YVals
}

func GetQuadTable(n int, x, y []float32) *QuadReg {
	Powers_Xi = 4
	Powers_YXi = 3

	xp := make([][]float32, Powers_Xi)
	xs := make([]float32, Powers_Xi)
	yp := make([][]float32, Powers_YXi)
	ys := make([]float32, Powers_YXi)

	for i := range 4 {
		xp[i] = make([]float32, n)
	}
	for i := range 3 {
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
	return []string{qr.GetCurve(solns)}, nil
}

//TODO: Eliminate cases like ... + -2x + -4.564x^2
func (*QuadReg) GetCurve(solns []float32) string {
	return fmt.Sprintf("y = %.3f + %.3fx + %.3fx^2", solns[0], solns[1], solns[2])
}

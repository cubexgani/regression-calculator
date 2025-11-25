package utils

import "fmt"

// Powers and sums indexed by power of x
// Here powers will have 4 rows: x^1, x^2, x^3, x^4. x^0 = 1, which isn't required.
// sums will contain their sums
type XVals struct {
	Num    int
	Powers [][]float32
	Sums   []float32
}

// Powers indexed by power of y
// Here powers will have 3 rows: y, xy, yx^2
// sums will contain their sums
type YVals struct {
	Num    int
	Powers [][]float32
	Sums   []float32
}

func GetTable(n int, x, y []float32) (XVals, YVals) {
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
	xv := XVals{n, xp, xs}
	yv := YVals{n, yp, ys}
	return xv, yv
}

//TODO: Eliminate cases like ... + -2x + -4.564x^2
func GetCurve(solns []float32) string {
	return fmt.Sprintf("y = %.3f + %.3fx + %.3fx^2", solns[0], solns[1], solns[2])
}

package utils

// Powers and sums indexed by power of x
// Here powers will have 4 rows: x^1, x^2, x^3, x^4. x^0 = 1, which isn't required.
// sums will contain their sums
type XVals struct {
	Powers [][]float32
	Sums   []float32
}

// Powers indexed by power of y
// Here powers will have 3 rows: y, xy, yx^2
// sums will contain their sums
type YVals struct {
	Powers [][]float32
	Sums   []float32
}

type Regression interface {
	Solve() ([]string, error)
	GetCurve([]float32) []string
}

func InitTable(n int, x, y []float32, regType string) (Regression, error) {
	var table Regression
	switch regType {
	case "quadratic":
		table = GetQuadTable(n, x, y)
	default:
		return nil, &InvalidRegTypeError{regType: regType}
	}
	return table, nil
}

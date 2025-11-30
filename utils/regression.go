package utils

var Powers_Xi, Powers_YXi int

// Powers and sums indexed by power of x
// Here powers will have rows like: x^1, x^2, x^3, x^4,... x^0 = 1, which isn't required.
// sums will contain their sums
type XVals struct {
	Powers [][]float32
	Sums   []float32
}

// Powers indexed by power of y
// Here powers will have rows like: y, xy, yx^2, yx^3,...
// sums will contain their sums
type YVals struct {
	Powers [][]float32
	Sums   []float32
}

type Regression interface {
	Solve() ([]string, error)
	GetCurve([]float32, rune) string
}

func InitTable(n int, x, y []float32, regType string) (Regression, error) {
	var table Regression
	switch regType {
	case "quadratic":
		table = GetQuadTable(n, x, y)
	case "linear":
		table = GetLinTable(n, x, y)
	default:
		return nil, &InvalidRegTypeError{regType: regType}
	}
	return table, nil
}

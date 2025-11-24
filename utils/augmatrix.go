package utils

import "fmt"

type AugMatrix struct {
	coeffs [][]float32
	vals   []float32
}

type InitError struct {
	coeffLen int
	valsLen  int
}

func (i InitError) Error() string {
	return fmt.Sprintf("Can't initialize augmented matrix as "+
		"length of coefficient matrix %d does not equal the length of "+
		"value vector %d", i.coeffLen, i.valsLen)
}

func (am AugMatrix) String() string {
	s := ""
	r := len(am.coeffs)
	c := len(am.coeffs[0])

	for i := range r {
		rs := ""
		for j := range c {
			rs = fmt.Sprintf("%s %.2f", rs, am.coeffs[i][j])
		}
		s = fmt.Sprintf("%s%s \t| %.2f\n", s, rs, am.vals[i])
	}

	return s
}

func (am *AugMatrix) Reduce() {
	r := len(am.coeffs)
	c := len(am.coeffs[0])

	if r == 0 || c == 0 {
		return
	}
	for i := range r - 1 {
		for j := i + 1; j < r; j++ {
			mult := am.coeffs[j][i] / am.coeffs[i][i]
			for k := range c {
				am.coeffs[j][k] -= mult * am.coeffs[i][k]
			}
			am.vals[j] -= mult * am.vals[i]
		}
	}
}

func (am *AugMatrix) GetSolutions() []float32 {
	solNum := len(am.coeffs[0])
	rows := len(am.coeffs)
	solns := make([]float32, solNum)

	for i := rows - 1; i >= 0; i-- {
		val := am.vals[i]
		for j := i; j < solNum-1; j++ {
			val -= am.coeffs[i][j+1] * solns[j+1]
		}
		solns[i] = val / am.coeffs[i][i]
	}

	return solns
}

func (am *AugMatrix) Rank() int {
	zerows := 0
	for i := range len(am.coeffs) {
		zeroCoeff := 1
		for j := range len(am.coeffs[0]) {
			if am.coeffs[i][j] != 0 {
				zeroCoeff = 0
			}
		}
		if zeroCoeff == 1 && am.vals[i] == 0 {
			zerows++
		}
	}
	return len(am.coeffs) - zerows
}

func MakeAugMat(coeff [][]float32, val []float32) (AugMatrix, error) {
	if len(coeff) != len(val) {
		return AugMatrix{}, &InitError{len(coeff), len(val)}
	}
	return AugMatrix{coeff, val}, nil
}

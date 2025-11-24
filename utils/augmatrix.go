package utils

import "fmt"

type AugMatrix struct {
	coeffs [][]int
	vals   []int
}

func (am AugMatrix) String() string {
	s := ""
	r := len(am.coeffs)
	c := len(am.coeffs[0])

	for i := range r {
		rs := ""
		for j := range c {
			rs = fmt.Sprintf("%s %d", rs, am.coeffs[i][j])
		}
		s = fmt.Sprintf("%s%s | %d\n", s, rs, am.vals[i])
	}

	return s
}

func (am *AugMatrix) reduce() {
	r := len(am.coeffs)
	c := len(am.coeffs[0])

	if r == 0 || c == 0 {
		return
	}

}

func Make_aug_mat(coeff [][]int, val []int) AugMatrix {
	return AugMatrix{coeff, val}
}

func test() {
	fmt.Println("Hello dawg")
}

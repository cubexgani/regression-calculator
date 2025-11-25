package utils

import "fmt"

type InitError struct {
	coeffLen int
	valsLen  int
}

type RankMismatchError struct {
	augRank   int
	coeffRank int
}

type HomogenousError struct {
	rank   int
	length int
}

func (ie InitError) Error() string {
	return fmt.Sprintf("Can't initialize augmented matrix as "+
		"length of coefficient matrix %d does not equal the length of "+
		"value vector %d", ie.coeffLen, ie.valsLen)
}

func (rme RankMismatchError) Error() string {
	return fmt.Sprintf("Rank of augmented matrix %d doesn't match "+
		"rank of coefficient matrix %d, implying no solution", rme.augRank, rme.coeffRank)
}

func (he HomogenousError) Error() string {
	return fmt.Sprintf("Rank of matrix %d doesn't match its length %d, "+
		"implying infinite solutions", he.rank, he.length)
}

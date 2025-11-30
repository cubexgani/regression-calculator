package utils

import "fmt"

// Error due to invalid regression type.
// I'm actually keeping this one aside for later use, might not even need it, who knows.
type InvalidRegTypeError struct {
	regType string
}

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

func (irte InvalidRegTypeError) Error() string {
	return fmt.Sprintf("Invalid regression type %s", irte.regType)
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

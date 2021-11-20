package fib

import (
	"math/big"
	"testing"
)

// 0,1,1,2,3,4,8,13,21,34,55,..
//func FibRec(n int) int {
func TestFibRec(t *testing.T) {
	data := []struct {
		/*
			n    uint
			want uint64
		*/
		n    int
		want int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := FibRec(d.n); got != d.want {
			t.Errorf("Bad FibRec value for n: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}

func TestBigFib1(t *testing.T) {
	f := BigFib(3)
	if f.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("Bad BigFib value for idx 3: %d", f)
	}
}

func TestBigTbl(t *testing.T) {
	data := []struct {
		n    int
		want int64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := BigFib(d.n); got.Int64() != d.want {
			t.Errorf("Bad FibRec value for n: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}

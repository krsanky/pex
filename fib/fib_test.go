package fib

import (
	"math/big"
	"testing"
)

func TestGrowFibVals(t *testing.T) {
	EnsureFibs(102)
}

func TestMemoFib1(t *testing.T) {
	f := MemoFib(3)
	if f.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("Bad MemoFib value for idx 3: %d", f)
	}
}

func TestMemoFibTbl(t *testing.T) {
	data := []struct {
		n    int
		want int64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := MemoFib(d.n); got.Int64() != d.want {

			t.Errorf("Bad MemoRec value for n: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}

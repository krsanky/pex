package fib

import (
	"math/big"
	"sync"
)

type FibsStore struct {
	sync.Mutex
	Vals [](*big.Int)
}

var Fib FibsStore

func init() {
	Fib.Lock()
	defer Fib.Unlock()
	n := 100
	Fib.Vals = make([](*big.Int), n)
	Fib.Vals[0] = big.NewInt(0)
	Fib.Vals[1] = big.NewInt(1)
	for i := 2; i < n; i++ {
		fib := &big.Int{}
		fib.Add(Fib.Vals[i-1], Fib.Vals[i-2])
		Fib.Vals[i] = fib
	}
}

func EnsureFibs(n int) {
	Fib.Lock()
	defer Fib.Unlock()

	if len(Fib.Vals) > n {
		return
	}

	for i := len(Fib.Vals); i <= n+1; i++ {
		//fmt.Printf("make new Fib.Vals i:%d len Fib.Vals:%d n:%d\n", i, len(Fib.Vals), n)
		fib := &big.Int{}
		fib.Add(Fib.Vals[i-1], Fib.Vals[i-2])
		Fib.Vals = append(Fib.Vals, fib)
	}
}

func MemoFib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(-1)
	}
	EnsureFibs(n)
	return Fib.Vals[n]
}

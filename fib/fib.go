package fib

import (
	"fmt"
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
	fmt.Printf("len Fib.Vals:%d\n", len(Fib.Vals))

	f100 := &big.Int{}
	f100.Add(Fib.Vals[100-1], Fib.Vals[100-2])
	Fib.Vals = append(Fib.Vals, f100)
	fmt.Printf("-len Fib.Vals:%d\n", len(Fib.Vals))

	// len Fib.Vals:100
	// -len Fib.Vals:101

}

func EnsureFibs(n int) {
	Fib.Lock()
	defer Fib.Unlock()

	if len(Fib.Vals) > n {
		return
	}

	for i := len(Fib.Vals); i <= n+1; i++ {
		fmt.Printf("make new Fib.Vals i:%d len Fib.Vals:%d n:%d\n", i, len(Fib.Vals), n)
		fib := &big.Int{}
		fib.Add(Fib.Vals[i-1], Fib.Vals[i-2])
		Fib.Vals = append(Fib.Vals, fib)
	}
}

func BigFib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(-1)
	}
	fibs := make([](*big.Int), n+1, n+2)
	if n < 2 {
		fibs = fibs[0:2]
	}
	fibs[0] = big.NewInt(0)
	fibs[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		fib := &big.Int{}
		fib.Add(fibs[i-1], fibs[i-2])
		fibs[i] = fib
	}
	return fibs[n]
}

func MemoFib(n int) *big.Int {
	EnsureFibs(n)
	return Fib.Vals[n]
}

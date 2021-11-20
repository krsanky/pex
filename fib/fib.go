package fib

import (
	"fmt"
	"math/big"
)

var BigFibs [](*big.Int)
var BigFibIdx int64

func FibRec(n int) int {
	if n < 0 {
		return -1
	}
	if n <= 1 {
		return n
	}
	return FibRec(n-1) + FibRec(n-2)
}

// -1 means error
func FibonacciLoop(n_ int) int64 {
	n := int64(n_)
	if n < 0 {
		return -1
	}
	f := make([]int64, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := int64(2); i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
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
		fmt.Printf("fib:%v\n", fib)
		fibs[i] = fib
	}
	return fibs[n]
}

/*
#
# Fast doubling Fibonacci algorithm (Python)
#
# (Public) Returns F(n).
def fibonacci(n):
	if n < 0:
		raise ValueError("Negative arguments not implemented")
	return _fib(n)[0]


# (Private) Returns the tuple (F(n), F(n+1)).
def _fib(n):
	if n == 0:
		return (0, 1)
	else:
		a, b = _fib(n // 2)
		c = a * (b * 2 - a)
		d = a * a + b * b
		if n % 2 == 0:
			return (c, d)
		else:
			return (d, c + d)
*/

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}

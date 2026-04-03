package main

import (
	"fmt"
)

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var nFact, nrFact int
	factorial(n, &nFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / nrFact
}

func combination(n, r int, hasil *int) {
	var nFact, rFact, nrFact int
	factorial(n, &nFact)
	factorial(r, &rFact)
	factorial(n-r, &nrFact)
	*hasil = nFact / (rFact * nrFact)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var pAC, cAC int
	permutation(a, c, &pAC)
	combination(a, c, &cAC)
	fmt.Println(pAC, cAC)

	var pBD, cBD int
	permutation(b, d, &pBD)
	combination(b, d, &cBD)
	fmt.Println(pBD, cBD)
}
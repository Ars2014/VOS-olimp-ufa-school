package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f1, err := os.Open("math.in")
	check(err)
	defer f1.Close()
	f2, err := os.Create("math.out")
	check(err)
	defer f2.Close()

	maxInt64 := big.NewInt(math.MaxInt64)

	for i := 0; i < 4; i++ {
		var first, second int64
		_, err = fmt.Fscan(f1, &first, &second)
		check(err)

		bigNum := new(big.Int)
		bigNum.Mul(big.NewInt(first), big.NewInt(second))
		if bigNum.Cmp(maxInt64) == 1 {
			fmt.Fprintln(f2, "YES")
		} else {
			fmt.Fprintln(f2, "NO")
		}
	}
}

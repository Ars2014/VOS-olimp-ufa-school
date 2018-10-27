package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func makeNum(x int) int64 {
	if x == 3 {
		return 1
	}
	n := float64(x - 2)
	return makeNum(x-1) + int64((2+(n-1))/2*n)
}

func getNum(n int64, nOfPeople int) int {
	prob := makeNum(nOfPeople)
	if n > prob {
		return getNum(n, nOfPeople+1)
	} else if n == prob {
		return nOfPeople
	} else {
		return 0
	}
}

func main() {
	var n int64
	f1, err := os.Open("meeting.in")
	check(err)
	defer f1.Close()
	_, err = fmt.Fscan(f1, &n)
	check(err)

	res := getNum(n, 3)

	f2, err := os.Create("meeting.out")
	check(err)
	defer f2.Close()
	if res != 0 {
		fmt.Fprintln(f2, res)
	} else {
		fmt.Fprintln(f2, "Err")
	}
}

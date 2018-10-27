package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f1, err := os.Open("binpattern.in")
	check(err)
	defer f1.Close()
	f2, err := os.Create("binpattern.out")
	check(err)
	defer f2.Close()

	for i := 0; i < 5; i++ {
		var num int
		_, err := fmt.Fscan(f1, &num)
		check(err)
		bin := fmt.Sprintf("%b", num)
		if strings.Contains(bin, "10101") {
			_, err = fmt.Fprintln(f2, "Pattern")
			check(err)
		} else {
			_, err = fmt.Fprintln(f2, "No")
			check(err)
		}
	}
}

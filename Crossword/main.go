package main

import (
	"bytes"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f1, err := os.Open("crossword.in")
	check(err)
	defer f1.Close()

	var first, second []byte
	_, err = fmt.Fscan(f1, &first, &second)
	check(err)

	var result int
	for _, letter := range first {
		n := bytes.Count(second, []byte{letter})
		result += n
	}

	f2, err := os.Create("crossword.out")
	check(err)
	defer f2.Close()
	_, err = fmt.Fprintln(f2, result)
	check(err)
}

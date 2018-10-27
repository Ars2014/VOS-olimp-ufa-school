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

func main() {
	var n int
	f1, err := os.Open("battleship.in")
	check(err)
	defer f1.Close()
	_, err = fmt.Fscanln(f1, &n)
	check(err)

	isFirst := true
	cells_fighted := make(map[string]bool)
	var (
		cells  []string
		result int
	)
	for i := 0; i < n; i++ {
		var (
			cell    string
			fighted int
		)
		_, err := fmt.Fscanf(f1, "%s %d\n", &cell, &fighted)
		check(err)
		if !isFirst && !cells_fighted[cell] {
			for _, secCell := range cells {
				if secCell == cell {
					result += 1
					cells_fighted[cell] = true
				}
			}
			cells = append(cells, cell)
		}
		if fighted == 0 {
			isFirst = !isFirst
		}
	}

	f2, err := os.Create("battleship.out")
	check(err)
	defer f2.Close()
	_, err = fmt.Fprint(f2, result)
	check(err)
}

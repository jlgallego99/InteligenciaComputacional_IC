package main

import (
	"bufio"
	"os"
	"strconv"
)

func ReadData(name string) (int, [][]int, [][]int, error) {
	var n int
	var A [][]int
	var B [][]int

	// QAP data file
	f, err := os.Open("./data/" + name)
	if err != nil {
		return 0, nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	// Read problem size (n)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	A = make([][]int, 0)
	B = make([][]int, 0)

	// Read matrix A
	row := make([]int, 0)
	r := 0
	for scanner.Scan() {
		d, err_conv := strconv.Atoi(scanner.Text())
		if err_conv != nil {
			return 0, nil, nil, err_conv
		}

		row = append(row, d)
		r++

		if r == n {
			A = append(A, row)
			row = make([]int, 0)
			r = 0
		}

		if len(A) == n {
			break
		}
	}

	// Read matrix B
	for scanner.Scan() {
		d, err_conv := strconv.Atoi(scanner.Text())
		if err_conv != nil {
			return 0, nil, nil, err_conv
		}

		row = append(row, d)
		r++

		if r == n {
			B = append(B, row)
			row = make([]int, 0)
			r = 0
		}
	}

	return n, A, B, nil
}

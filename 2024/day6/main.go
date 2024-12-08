package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read_file_to_lines(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		print(err)
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data
}

func data_to_matrix(data []string) [][]string {
	matrix := make([][]string, len(data))
	for ix, line := range data {
		line_slice := strings.Split(line, "")
		matrix[ix] = make([]string, len(line_slice))
		for jx := 0; jx < len(line_slice); jx++ {
			matrix[ix][jx] = line_slice[jx]
		}
	}
	return matrix
}

func find_start(mappa [][]string) (int, int) {
	rows := len(mappa)
	cols := len(mappa[0])
	for ix := 0; ix < rows; ix++ {
		for jx := 0; jx < cols; jx++ {
			if mappa[ix][jx] == "^" {
				return ix, jx
			}
		}
	}
	return 0, 0
}

func find_next_step_cand(mappa [][]string, ix int, jx int) (int, int) {
	guard_type := mappa[ix][jx]
	switch guard_type {
	case "^":
		ix--
	case ">":
		jx++
	case "v":
		ix++
	case "<":
		jx--
	}
	return ix, jx
}

func take_a_walk(mappa [][]string) {
	rotate := map[string]string{"<": "^", "^": ">", ">": "v", "v": "<"}
	ix, jx := find_start(mappa)
	guard_type := "^"
	rows := len(mappa)
	cols := len(mappa[0])
	ixn := 0
	jxn := 0
	safety := 0
	for {
		ixn, jxn = find_next_step_cand(mappa, ix, jx)
		if ixn < 0 || jxn < 0 || ixn == rows || jxn == cols { //check OOB
			mappa[ix][jx] = "X"
			fmt.Println("OUT")
			break
		} else if mappa[ixn][jxn] == "#" { //Check hinder
			guard_type = rotate[guard_type]
			mappa[ix][jx] = guard_type
		} else {
			mappa[ixn][jxn] = guard_type
			mappa[ix][jx] = "X"
			ix = ixn
			jx = jxn
		}
		safety++
		if safety > 1<<20 {
			break
		}
	}
}

func count_X(mappa [][]string) int {
	rows := len(mappa)
	cols := len(mappa[0])
	tot := 0
	for ix := 0; ix < rows; ix++ {
		for jx := 0; jx < cols; jx++ {
			if mappa[ix][jx] == "X" {
				tot++
			}
		}
	}
	return tot
}
func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day6.txt"
	data := read_file_to_lines(filename)
	matrix := data_to_matrix(data)
	take_a_walk(matrix)
	for _, line := range matrix {
		fmt.Println(line)
	}
	fmt.Println(count_X(matrix))
}

func main() {
	task1()
}

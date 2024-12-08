package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type turn struct {
	ix             int
	jx             int
	dir_after_turn string
}

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

func find_next_step_cand2(guard_type string, ix int, jx int) (int, int) {
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
			//fmt.Println("OUT")
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
	//for _, line := range matrix {
	//fmt.Println(line)
	//}
	fmt.Println(count_X(matrix))
}

func take_a_walk2(mappa [][]string) int {
	rotate := map[string]string{"<": "^", "^": ">", ">": "v", "v": "<"}
	step_type := map[string]string{"<": "-", "^": "|", ">": "-", "v": "|", "rotate": "+"}
	var turns []turn
	ix, jx := find_start(mappa)
	guard_type := "^"
	rows := len(mappa)
	cols := len(mappa[0])
	ixn := 0
	jxn := 0
	safety := 0
	rotate_last_step := false
	for {
		ixn, jxn = find_next_step_cand2(guard_type, ix, jx)
		if ixn < 0 || jxn < 0 || ixn == rows || jxn == cols { //check OOB
			mappa[ix][jx] = step_type[guard_type]
			break
		} else if mappa[ixn][jxn] == "#" || mappa[ixn][jxn] == "O" { //Check hinder
			guard_type = rotate[guard_type]
			rotate_last_step = true
			this_turn := turn{ix: ix, jx: jx, dir_after_turn: guard_type}
			n_turns := len(turns)
			if n_turns > 4 {
				for _, a_turn := range turns {
					if this_turn == a_turn {
						//fmt.Println("LOOP_FOUND")
						return 1
					}
				}
			}
			turns = append(turns, this_turn)
		} else {
			if rotate_last_step {
				rotate_last_step = false
				mappa[ix][jx] = step_type["rotate"]
			} else {
				mappa[ix][jx] = step_type[guard_type]
			}
			ix = ixn
			jx = jxn
		}
		safety++
		if safety > 1<<20 {
			break
		}
	}
	return 0
}

func addHinder(ix, jx int, data []string) [][]string {
	mappa := data_to_matrix(data)
	mappa[ix][jx] = "O"
	//fmt.Println(ix, jx)
	//fmt.Println("OH NO! THE SWEDISH HINDER!")
	return mappa
}

func task2() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day6.txt"
	data := read_file_to_lines(filename)
	mappa := data_to_matrix(data)
	rows := len(mappa)
	cols := len(mappa[0])
	ix_start, jx_start := find_start(mappa)
	sum := 0
	n_hinders := 0
	for ix := 0; ix < rows; ix++ {
		for jx := 0; jx < cols; jx++ {
			if (ix == ix_start) && (jx == jx_start) {
			} else {
				n_hinders++
				mappa = addHinder(ix, jx, data)
				res := take_a_walk2(mappa)
				if res > 0 {
					//fmt.Println(ix+1, jx+1)
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(n_hinders)
	//for _, line := range mappa {
	//	fmt.Println(line)
	//}
}
func main() {
	//	task1()
	task2()
}

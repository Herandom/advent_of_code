package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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
	var lines []string
	ix := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		ix++
	}
	return lines
}

func find_xmassamx(input string) int {
	re := regexp.MustCompile("XMAS")
	re2 := regexp.MustCompile("SAMX")
	xmas := len(re.FindAllString(input, -1))
	samx := len(re2.FindAllString(input, -1))
	return xmas + samx

}

func to_matrix(input []string) [][]string {
	matrix := make([][]string, len(input))
	for ix, line := range input {
		line_arr := strings.Split(line, "")
		matrix[ix] = make([]string, len(line_arr))
		for jx := 0; jx < len(line_arr); jx++ {
			matrix[ix][jx] = line_arr[jx]
		}
	}
	return matrix
}
func transpose(matrix [][]string) [][]string {
	row := len(matrix)
	col := len(matrix[0])
	transp := make([][]string, row)
	for ix := 0; ix < row; ix++ {
		transp[ix] = make([]string, col)
		for jx := 0; jx < col; jx++ {
			transp[ix][jx] = matrix[jx][ix]
		}
	}
	return transp
}

func rotate(matrix [][]string) [][]string {
	dim := len(matrix)
	rotated := make([][]string, 2*dim-1)
	diaglen := 0
	for kx := 0; kx < 2*dim-1; kx++ {
		if kx < dim {
			diaglen++
		} else {
			diaglen--
		}
		rotated[kx] = make([]string, diaglen)
		dx := 0
		for jx := 0; jx <= kx; jx++ {
			ix := kx - jx
			if ix < dim && jx < dim {
				rotated[kx][dx] = matrix[ix][jx]
				dx++
			}
		}
	}

	return rotated
}

func reverse_matrix_rows(matrix [][]string) [][]string {
	reversed := make([][]string, len(matrix))
	for ix, row := range matrix {
		clone := slices.Clone(row)
		slices.Reverse(clone)
		reversed[ix] = clone
	}

	return reversed
}
func search_matrix(matrix [][]string) int {
	match := 0
	for _, row_slice := range matrix {
		row := strings.Join(row_slice, "")
		match += find_xmassamx(row)
	}
	return match
}

func search_x_mas(matrix [][]string) int {
	rows := len(matrix)
	cols := len(matrix[0])
	re := regexp.MustCompile("MAS|SAM")
	match := 0
	for ix := 1; ix < rows-1; ix++ {
		for jx := 1; jx < cols-1; jx++ {
			if matrix[ix][jx] == "A" {
				diag1 := []string{matrix[ix-1][jx-1], matrix[ix][jx], matrix[ix+1][jx+1]}
				diag2 := []string{matrix[ix+1][jx-1], matrix[ix][jx], matrix[ix-1][jx+1]}
				diag1_str := strings.Join(diag1, "")
				diag2_str := strings.Join(diag2, "")
				if re.MatchString(diag1_str) && re.MatchString(diag2_str) {
					match++
				}
			}
		}
	}
	return match
}

func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day4.txt"
	lines := read_file_to_lines(filename)
	matches := 0
	matrix := to_matrix(lines)
	transp := transpose(matrix)
	reverse := reverse_matrix_rows(matrix)
	diag_1 := rotate(matrix)
	diag_2 := rotate(reverse)
	matches += search_matrix(matrix)
	matches += search_matrix(transp)
	matches += search_matrix(diag_1)
	matches += search_matrix(diag_2)
	fmt.Println(matches)
}

func task2() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day4.txt"
	lines := read_file_to_lines(filename)
	matches := 0
	matrix := to_matrix(lines)
	matches = search_x_mas(matrix)
	fmt.Println(matches)
}

func main() {
	task1()
	task2()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func split_input(data []string) map[int64][]int64 {
	data_map := map[int64][]int64{}
	for _, line := range data {
		split := strings.Split(line, ": ")
		ans, _ := strconv.Atoi(split[0])
		inpstr := strings.Split(split[1], " ")
		ints := make([]int64, len(inpstr))
		for ix, elem := range inpstr {
			temp, _ := (strconv.Atoi(elem))
			ints[ix] = int64(temp)
		}
		data_map[int64(ans)] = ints
	}
	return data_map
}

func generate_permutations(operators []string, all_operators *[]string, ix int, length int) {
	if ix == length {
		opers := strings.Join(operators, ",")

		*all_operators = append(*all_operators, opers)
		return
	}

	operators[ix] = "+"
	generate_permutations(operators, all_operators, ix+1, length)

	operators[ix] = "x"
	generate_permutations(operators, all_operators, ix+1, length)

	//operators[ix] = "-"
	//generate_permutations(operators, all_operators, ix+1, length)

	//operators[ix] = "/"
	//generate_permutations(operators, all_operators, ix+1, length)
}

func check_valid_datas(data []string) int64 {
	permutation_map := map[int][][]string{}
	is_successful := 0
	var summus_totalitus int64 = 0
	niter := 1
	for _, data_line := range data {
		split := strings.Split(data_line, ": ")
		ans_in, _ := strconv.Atoi(split[0])
		ans := int64(ans_in)
		inpstr := strings.Split(split[1], " ")
		input := make([]int64, len(inpstr))
		for ix, elem := range inpstr {
			temp, _ := (strconv.Atoi(elem))
			input[ix] = int64(temp)
		}

		niter++
		n_ops := len(input) - 1
		_, ok := permutation_map[n_ops]
		if !ok {
			operators := make([]string, n_ops)
			all_operators := make([]string, 0)
			generate_permutations(operators, &all_operators, 0, n_ops)
			opers := make([][]string, len(all_operators))
			for ix, line := range all_operators {
				line_slice := strings.Split(line, ",")
				opers[ix] = make([]string, len(line_slice))
				for jx, elem := range line_slice {
					opers[ix][jx] = elem
				}
			}
			permutation_map[n_ops] = opers
		}
		opmat := permutation_map[n_ops]
		for _, op_set := range opmat {
			res := input[0]
			for ix, op := range op_set {
				res = use_operation(res, input[ix+1], op)
			}
			if res == ans {
				is_successful++
				summus_totalitus += ans
				break
			}
		}
	}
	fmt.Printf("N successful: %v\n", is_successful)
	return summus_totalitus
}

func generate_permutations2(operators []string, all_operators *[]string, ix int, length int) {
	if ix == length {
		opers := strings.Join(operators, ",")

		*all_operators = append(*all_operators, opers)
		return
	}

	operators[ix] = "+"
	generate_permutations2(operators, all_operators, ix+1, length)

	operators[ix] = "x"
	generate_permutations2(operators, all_operators, ix+1, length)

	operators[ix] = "||"
	generate_permutations2(operators, all_operators, ix+1, length)

}

func use_operation(a, b int64, oper string) int64 {
	switch oper {
	case "+":
		return a + b
	case "x":
		return a * b
	case "||":
		a_str := strconv.Itoa(int(a))
		b_str := strconv.Itoa(int(b))
		c_t, _ := strconv.Atoi(a_str + b_str)
		c := int64(c_t)
		//fmt.Println(a, a_str, b, b_str, c)
		return c
	}
	return 0
}

func check_valid_datas2(data []string) int64 {
	permutation_map := map[int][][]string{}
	is_successful := 0
	var summus_totalitus int64 = 0
	niter := 1
	for _, data_line := range data {
		split := strings.Split(data_line, ": ")
		ans_in, _ := strconv.Atoi(split[0])
		ans := int64(ans_in)
		inpstr := strings.Split(split[1], " ")
		input := make([]int64, len(inpstr))
		for ix, elem := range inpstr {
			temp, _ := (strconv.Atoi(elem))
			input[ix] = int64(temp)
		}

		niter++
		n_ops := len(input) - 1
		_, ok := permutation_map[n_ops]
		if !ok {
			operators := make([]string, n_ops)
			all_operators := make([]string, 0)
			generate_permutations2(operators, &all_operators, 0, n_ops)
			opers := make([][]string, len(all_operators))
			for ix, line := range all_operators {
				line_slice := strings.Split(line, ",")
				opers[ix] = make([]string, len(line_slice))
				for jx, elem := range line_slice {
					opers[ix][jx] = elem
				}
			}
			permutation_map[n_ops] = opers
		}
		opmat := permutation_map[n_ops]
		for _, op_set := range opmat {
			res := input[0]
			for ix, op := range op_set {
				res = use_operation(res, input[ix+1], op)
			}
			if res == ans {
				is_successful++
				summus_totalitus += ans
				break
			}
		}
	}
	fmt.Printf("N successful: %v\n", is_successful)
	return summus_totalitus
}

func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day7.txt"

	data := read_file_to_lines(filename)
	fmt.Println(check_valid_datas(data))
}

func task2() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day7.txt"

	data := read_file_to_lines(filename)
	fmt.Println(check_valid_datas2(data))
}
func main() {
	task1()
	fmt.Print("\n\n")
	task2()
}

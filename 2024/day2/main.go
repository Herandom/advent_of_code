package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day2.txt"
	allowed := []int{-1, -2, -3, 1, 2, 3}
	var diff int
	count := 0
	ok := true
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var codes []int
		for _, item := range line {
			code_int, _ := strconv.Atoi(item)
			codes = append(codes, code_int)
		}
		var diffs []int
		ok = true
		for ix := 0; ix < len(codes)-1; ix++ {
			diff = codes[ix] - codes[ix+1]
			diffs = append(diffs, diff)
			if !slices.Contains(allowed, diff) {
				ok = false
				break
			}

		}
		if ok && (slices.Min(diffs)*slices.Max(diffs) > 0) {

			count += 1
		}
	}
	defer file.Close()
	fmt.Print(count, "\n")
}
func main() {
	task1()
	fmt.Print("hrj<2")

}

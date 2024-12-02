package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"strconv"
	"strings"
)

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func task1() {
	var lefts, rights []int
	//	filename := "test_data.txt"

	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		left, _ := strconv.Atoi(line[0])
		right, _ := strconv.Atoi(line[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
		//	fmt.Printf("%v, %v \n", lefts, rights)

	}
	defer file.Close()
	sort.Ints(lefts)
	sort.Ints(rights)
	diff := 0
	for ix := 0; ix < len(lefts); ix++ {
		diff += abs(lefts[ix] - rights[ix])
	}
	fmt.Println(diff)
}

func task2() {
	var lefts, rights []int
	//filename := "test_data.txt"

	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		left, _ := strconv.Atoi(line[0])
		right, _ := strconv.Atoi(line[1])
		lefts = append(lefts, left)
		rights = append(rights, right)
		//	fmt.Printf("%v, %v \n", lefts, rights)

	}
	defer file.Close()
	sort.Ints(lefts)
	sort.Ints(rights)
	my_map := make(map[int]int)
	val_count := make(map[int]int)
	for _, val := range lefts {
		my_map[val] = 0
		val_count[val] += 1
	}
	for _, val := range rights {
		_, ok := my_map[val]
		if ok {
			my_map[val] += 1
		}
	}
	sim := 0
	for val, count := range my_map {
		sim += val * count * val_count[val]
	}
	fmt.Println(sim)
}
func main() {
	// (	task1()
	task2()
}

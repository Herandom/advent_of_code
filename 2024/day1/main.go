package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lefts, rights []int
	filename := "test_data.txt"
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
		fmt.Printf("%v, %v \n", lefts, rights)
	}
}

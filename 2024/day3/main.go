package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mult(fac string) int {
	facs := strings.Split(fac, ",")
	a, _ := strconv.Atoi(facs[0])
	b, _ := strconv.Atoi(facs[1])
	return a * b
}
func task1() {
	filename := "test_data.txt"
	//filename = "/home/herandom/projects/advent_of_code/inputs/2024/day3.txt"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var input string

	for scanner.Scan() {
		line := scanner.Text()
		input = input + line
		//fmt.Print(input)
	}
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re2 := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	match1 := strings.Join(re.FindAllString(input, -1), "")
	match2 := re2.FindAllString(match1, -1)
	fmt.Printf("%q\n", match1)
	fmt.Printf("%q\n", match2)
	sum := 0
	for _, item := range match2 {
		sum += mult(item)
	}
	fmt.Println(sum)
}

func main() {
	task1()
}

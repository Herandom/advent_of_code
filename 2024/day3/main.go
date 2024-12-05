package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_file_to_str(filename string) string {
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
	}
	return input
}
func mult(fac string) int {
	facs := strings.Split(fac, ",")
	a, _ := strconv.Atoi(facs[0])
	b, _ := strconv.Atoi(facs[1])
	return a * b
}

func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day3.txt"
	input := read_file_to_str(filename)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re2 := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	match1 := strings.Join(re.FindAllString(input, -1), "")
	match2 := re2.FindAllString(match1, -1)
	//	fmt.Printf("%q\n", match1)
	//	fmt.Printf("%q\n", match2)
	sum := 0
	for _, item := range match2 {
		sum += mult(item)
	}
	fmt.Println(sum)
}

func task2() {
	filename := "test_data2.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day3.txt"
	input := read_file_to_str(filename)
	re_dodont := regexp.MustCompile(`don't\(\).*?do\(\)`)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re_dodont_fin := regexp.MustCompile(`do\(\).*?don't\(\)`)
	re2 := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	matchslice := re_dodont.Split(input, -1)
	matchslice[len(matchslice)-1] = re_dodont_fin.FindString(matchslice[len(matchslice)-1])
	match := strings.Join(matchslice, "")
	match1 := strings.Join(re.FindAllString(match, -1), "")
	match2 := re2.FindAllString(match1, -1)
	//	fmt.Printf("%q\n", input)
	fmt.Printf("%q\n", matchslice[len(matchslice)-1])
	fmt.Printf("%q\n", match)
	//fmt.Printf("%q\n", match1)
	//	fmt.Printf("%q\n", match2)
	sum := 0
	for _, item := range match2 {
		sum += mult(item)
	}
	fmt.Println(sum)
}
func main() {
	task1()
	task2()
}

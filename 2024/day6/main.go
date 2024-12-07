package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_file_to_lines(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		print(err)
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var rules []string
	var data []string
	ix := 0
	read_rules := true
	for scanner.Scan() {
		line := scanner.Text()
		ix++
		if line == "" {
			fmt.Println("TOMELITOM")
			read_rules = false
		} else if read_rules {
			rules = append(rules, line)
		} else {
			data = append(data, line)
		}
	}
	return rules, data
}


func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day5.txt"

}


}
func main() {
	task1()
}

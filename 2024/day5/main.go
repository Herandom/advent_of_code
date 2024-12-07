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

func make_rule_invert_map_ints(rules []string) map[int][]int {
	inv_rule := map[int][]int{}
	for _, rulestr := range rules {
		rule_strs := strings.Split(rulestr, "|")
		a, _ := strconv.Atoi(rule_strs[0])
		b, _ := strconv.Atoi(rule_strs[1])
		_, ok := inv_rule[b]
		if ok {
			inv_rule[b] = append(inv_rule[b], a)
		} else {
			inv_rule[b] = []int{a}
		}
	}
	return inv_rule
}

func make_rule_invert_map(rules []string) map[string][]string {
	inv_rule := map[string][]string{}
	for _, rulestr := range rules {
		rule_strs := strings.Split(rulestr, "|")
		a := rule_strs[0]
		b := rule_strs[1]
		_, ok := inv_rule[b]
		if ok {
			inv_rule[b] = append(inv_rule[b], a)
		} else {
			inv_rule[b] = []string{a}
		}
	}
	return inv_rule
}

func rules_fail(rules []string, pages []string) bool {
	inter := make([]string, 0)
	pages_map := make(map[string]struct{}, 0)
	for _, page := range pages {
		pages_map[page] = struct{}{}
	}
	for _, rule := range rules {
		if _, found := pages_map[rule]; found {
			inter = append(inter, rule)
		}
	}
	return len(inter) != 0
}

func check_data(data []string, rules map[string][]string) []string {
	good_data := make([]string, 0)
	is_good := true
	for _, line := range data {
		is_good = true
		pages := strings.Split(line, ",")
		for ix, page := range pages {
			if rules_fail(rules[page], pages[ix:]) {
				is_good = false
				break
			}
		}
		if is_good {
			good_data = append(good_data, line)
		}
	}
	return good_data
}

func check_data_bad(data []string, rules map[string][]string) []string {
	bad_data := make([]string, 0)
	is_bad := true
	for _, line := range data {
		is_bad = false
		pages := strings.Split(line, ",")
		for ix, page := range pages {
			if rules_fail(rules[page], pages[ix:]) {
				is_bad = true
				break
			}
		}
		if is_bad {
			bad_data = append(bad_data, line)
		}
	}
	return bad_data
}
func sum_middle(data []string) int {
	sum := 0
	for _, line := range data {
		pages := strings.Split(line, ",")
		midelem := (len(pages) - 1) / 2
		mid, _ := strconv.Atoi(pages[midelem])
		sum += mid
	}
	return sum
}

func task1() {
	filename := "test_data.txt"
	filename = "/home/herandom/projects/advent_of_code/inputs/2024/day5.txt"
	rules, data := read_file_to_lines(filename)
	forbidden := make_rule_invert_map(rules)
	good_data := check_data(data, forbidden)
	sum := sum_middle(good_data)
	fmt.Println(sum)

}

func task2() {
	filename := "test_data.txt"
	//filename = "/home/herandom/projects/advent_of_code/inputs/2024/day5.txt"
	rules, data := read_file_to_lines(filename)
	forbidden := make_rule_invert_map(rules)
	bad_data := check_data_bad(data, forbidden)
	//sum := sum_middle(good_data)
	fmt.Println(bad_data)

}
func main() {
	//task1()
	task2()
}

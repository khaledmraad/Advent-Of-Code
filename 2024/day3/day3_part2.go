package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func res(match string) int64 {
	var res int64

	// fmt.Println(match[4:strings.Index(match, ",")])
	// fmt.Println(match[strings.Index(match, ",")+1 : len(match)-1])

	a, _ := strconv.Atoi(match[4:strings.Index(match, ",")])
	b, _ := strconv.Atoi(match[strings.Index(match, ",")+1 : len(match)-1])

	res += (int64)(a * b)

	return res
}

func part2() {

	body, err := os.Open("./input3_part2")
	if err != nil {
		fmt.Println(err)
	}

	defer body.Close()

	scanner := bufio.NewScanner(body)

	// reader := bufio.NewReader(strings.NewReader(string(body)))

	// var file_content [][]int

	var result []string
	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)
		if err != nil {
			break
		}
	}
	var final_res int64
	var test bool = true
	for _, i := range result {
		var re = regexp.MustCompile(`(?m)mul\((-?\d+),(-?\d+)\)|do\(\)|don't\(\)`)
		var str = i
		matches := re.FindAllString(str, -1)

		fmt.Println(matches)
		for i := 0; i < len(matches); i++ {
			if !strings.Contains(matches[i], "do") && test {
				final_res += res(matches[i])
			}
			if matches[i] == "do()" {
				test = true
			}
			if matches[i] == "don't()" {
				test = false
			}

		}

	}
	println(final_res)
}

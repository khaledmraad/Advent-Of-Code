package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {

	body, err := os.Open("./input3_part1")
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
	var res int64
	for _, i := range result {
		var re = regexp.MustCompile(`(?m)mul\((-?\d+),(-?\d+)\)`)
		var str = i

		for _, match := range re.FindAllString(str, -1) {

			fmt.Println(match)
			fmt.Println(match[4:strings.Index(match, ",")])
			fmt.Println(match[strings.Index(match, ",")+1 : len(match)-1])

			a, erra := strconv.Atoi(match[4:strings.Index(match, ",")])
			b, errb := strconv.Atoi(match[strings.Index(match, ",")+1 : len(match)-1])

			res += (int64)(a * b)

			if erra != nil || errb != nil {
				break
			}
		}

	}
	println(res)

}

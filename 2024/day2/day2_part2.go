package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func safe(v []int) bool {
	var diff []int
	for i := 0; i < len(v)-1; i++ {
		diff = append(diff, v[i]-v[i+1])
	}
	// fmt.Println(diff)
	for i := 0; i < len(diff); i++ {
		if i == 0 && (math.Abs(float64(diff[i])) > 3 || diff[i] == 0) {
			return false
		}
		if i == 0 {
			continue
		}
		if (diff[i] > 0 && diff[i-1] < 0) || (diff[i] < 0 && diff[i-1] > 0) || (math.Abs(float64(diff[i])) > 3) || (diff[i] == 0) {
			return false
		}
		if i == len(diff)-1 {
			return true
		}
	}
	return false
}

func part2() {

	body, err := os.Open("./input2_part2")
	if err != nil {
		fmt.Println(err)
	}

	defer body.Close()

	scanner := bufio.NewScanner(body)

	// reader := bufio.NewReader(strings.NewReader(string(body)))

	var file_content [][]int

	for scanner.Scan() {
		// line, err := reader.ReadString('\n')
		var result []int
		line := scanner.Text()
		// fmt.Println(line)
		for _, v := range strings.Split(line, " ") {
			// if x == len(strings.Split(line, " "))-1 {
			// 	v = v[:len(v)-1]
			// 	fmt.Println(v)
			// }
			i, err1 := strconv.Atoi(v)
			result = append(result, i)
			if err1 != nil {
				break
			}
		}

		file_content = append(file_content, result)

		if err != nil {
			break
		}
	}

	// fmt.Println(file_content)
	var count int = 0
	for _, v := range file_content {

		if safe(v) {
			count++
			continue
		} else {
			for i := 0; i < len(v); i++ {
				var test []int = append(append([]int{}, v[:i]...), v[i+1:]...)
				if safe(test) {
					count++
					break
				}
			}
		}
	}

	fmt.Println(count)

}

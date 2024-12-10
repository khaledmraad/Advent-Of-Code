package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var left_slice = []int{}
	var right_slice = []int{}

	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
		result := strings.Split(line, "   ")
		left, err1 := strconv.Atoi(result[0])
		right, err2 := strconv.Atoi(result[1][:len(result[1])-1])
		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}
		left_slice = append(left_slice, left)
		right_slice = append(right_slice, right)
	}

	sort.Ints(left_slice)
	sort.Ints(right_slice)

	var res float64

	for i := 0; i < len(left_slice); i++ {
		res += math.Abs(float64(left_slice[i] - right_slice[i]))
	}
	fmt.Printf("%.20f\n", res)
}

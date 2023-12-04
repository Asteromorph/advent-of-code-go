package p4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetOverlapAtAll() {
	input, _ := os.Open("./pkg/p4/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	sum := 0

	for sc.Scan() {
		tokens := strings.Split(sc.Text(), ",")
		first := strings.Split(tokens[0], "-")
		second := strings.Split(tokens[1], "-")

		startFirst, _ := strconv.Atoi(first[0])
		endFirst, _ := strconv.Atoi(first[1])
		startSecond, _ := strconv.Atoi(second[0])
		endSecond, _ := strconv.Atoi(second[1])

		if startSecond <= endFirst && endSecond >= startFirst || startFirst <= endSecond && endFirst >= startSecond {
			sum += 1
		}
	}
	fmt.Println(sum)
}

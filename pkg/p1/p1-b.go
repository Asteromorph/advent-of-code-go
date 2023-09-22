package p1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetMaxCalories2() int {
	input, _ := os.Open("input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)

	top1 := 0
	top2 := 0
	top3 := 0
	tmp := 0
	curr := 0
	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())
		curr += snack

		if err != nil {
			if curr > top3 {
				top3 = curr
			}
			if top3 > top2 {
				tmp = top2
				top2 = top3
				top3 = tmp
			}
			if top2 > top1 {
				tmp = top1
				top1 = top2
				top2 = tmp
			}
			fmt.Printf("curr: %v %v %v %v \n", top1, top2, top3, curr)
			curr = 0
		}
	}
	curr = top1 + top2 + top3
	return curr
}

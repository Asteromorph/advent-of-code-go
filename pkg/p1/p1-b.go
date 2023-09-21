package p1

import (
	"bufio"
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
    curr := 0
    for sc.Scan() {
	snack, err := strconv.Atoi(sc.Text())
	curr += snack

	if err != nil {
	    if curr > top3 {
		top3 = curr 
	    }	    
	    if top3 > top2 {
		top2, top3 = top3, top2
	    }
	    if top2 > top1 {
		top2, top1 = top1, top2
	    }
	}
	curr = 0
    }
    return curr
}

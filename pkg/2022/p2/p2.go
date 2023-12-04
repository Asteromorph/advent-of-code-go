package p2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetTotalPoints() {
	input, _ := os.Open("./pkg/p2/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	total := 0
	for sc.Scan() {
		if sc.Err() == nil {
			tokens := strings.Split(sc.Text(), " ")
			switch tokens[1] {
			case "X":
				GetResultX(tokens[0], &total)
			case "Y":
				GetResultY(tokens[0], &total)
			case "Z":
				GetResultZ(tokens[0], &total)
			}
		}
        fmt.Printf("total: %d\n",total)
	}
}

func GetResultX(other string, total *int) {
	*total += 1
	switch other {
	case "A":
		*total += 3
	case "B":
		*total += 0
	case "C":
		*total += 6
	}
    fmt.Printf("X: %v\n", *total)

}
func GetResultY(other string, total *int) {
	*total += 2
	switch other {
	case "A":
		*total += 6
	case "B":
		*total += 3
	case "C":
		*total += 0
	}
}

func GetResultZ(other string, total *int) {
	*total += 3
	switch other {
	case "A":
		*total += 0
	case "B":
		*total += 6
	case "C":
		*total += 3
	}

}

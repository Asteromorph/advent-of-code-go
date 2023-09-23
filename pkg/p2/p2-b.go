package p2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetTotalPoints2() {
	input, _ := os.Open("./pkg/p2/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	total := 0
	for sc.Scan() {
		if sc.Err() == nil {
			tokens := strings.Split(sc.Text(), " ")
			switch tokens[1] {
			case "X":
				getLosingResult(tokens[0], &total)
			case "Y":
				getDrawResult(tokens[0], &total)
			case "Z":
				getWinResult(tokens[0], &total)
			}
		}
        fmt.Printf("total: %d\n",total)
	}
}

func getLosingResult(other string, total *int) {
	switch other {
	case "A":
		*total += 3
	case "B":
		*total += 1
	case "C":
		*total += 2
	}
}
func getDrawResult(other string, total *int) {
	*total += 3
	switch other {
	case "A":
		*total += 1
	case "B":
		*total += 2
	case "C":
		*total += 3
	}
}

func getWinResult(other string, total *int) {
	*total += 6
	switch other {
	case "A":
		*total += 2
	case "B":
		*total += 3
	case "C":
		*total += 1
	}

}

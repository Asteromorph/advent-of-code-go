package d15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	focalLength int
	operation rune
}

func Part2() {
	input, _ := os.Open("./pkg/2023/d15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var str string
	for sc.Scan() {
		str += sc.Text()
	}
	boxes := map[int][]Lens{}
	for _, lens := range parseLine(str) {
		placeLens(lens, boxes)
	}
	fmt.Println("res:", getFocusingPower(boxes))
}

func parseLens(str string) (lens Lens) {
	pos := -1
	if pos = strings.IndexRune(str, '='); pos != -1 {
		lens.operation = '='
	} else {
		pos = strings.IndexRune(str, '-')
		lens.operation = '-'
	}
	lens.label = str[:pos]
	lens.focalLength, _ = strconv.Atoi(str[len(str) - 1:])
	return
}

func placeLens(lensStr string, boxes map[int][]Lens) {
	lens := parseLens(lensStr)
	boxNumber := hash(lens.label)
	if val, ok := boxes[boxNumber]; ok {
		if lens.operation == '-' {
			for i, v := range val {
				if v.label == lens.label {
					val = append(val[:i], val[i+1:]...)
					boxes[boxNumber] = val
				}
			}
		} else {
			for i, v := range val {
				if v.label == lens.label {
					val[i] = lens
					return
				}
			}
			val = append(val, lens)
			boxes[boxNumber] = val
		}
	} else {
		if lens.operation == '=' {
			lenses := []Lens{}
			lenses = append(lenses, lens)
			boxes[boxNumber] = lenses
		}
	}
}

func getFocusingPower(boxes map[int][]Lens) (res int){
	for i, v := range boxes {
		if len(v) != 0 {
			for li, lv := range v {
				res += (i + 1) * (li + 1) * lv.focalLength
			}
		}
	}
	return 
}

func parseLine(str string) []string {
	return strings.Split(str, ",")
}

func hash(str string) int {
	res := 0
	for _, v := range str {
		res += int(v)
		res *= 17
		res = res % 256
	}
	return res
}

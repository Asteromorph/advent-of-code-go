package p10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DrawCRT() {
	input, _ := os.Open("./pkg/p10/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	cycle, x := 0, 1 
	for sc.Scan() {
		opts := strings.Fields(sc.Text())		
		draw(&cycle, &x)
		if opts[0] == "addx" {
			draw(&cycle, &x)
			value, _ := strconv.Atoi(opts[1])
			x += value
		}
	}
}

func draw(cycle, x *int) {
	if *cycle % 40 == 0 && *cycle < 221 {
		fmt.Println()
	}
	if *cycle%40 - 1 == *x || *cycle%40 == *x || *cycle%40 + 1 == *x {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	*cycle++
}

package p10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TotalSignals() {
	input, _ := os.Open("./pkg/p10/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	cycle, x, total := 0, 1 ,0
	for sc.Scan() {
		opts := strings.Fields(sc.Text())
		calculate(&cycle, &x, &total)
		if opts[0] == "addx" {
			calculate(&cycle, &x, &total)
			value, _ := strconv.Atoi(opts[1])
			x += value
		}
	}
	fmt.Println(total)
}

func calculate(cycle, x, total *int) {
	*cycle++
	if (*cycle - 20)%40 == 0 && *cycle <= 221 {
		*total += *x * *cycle
	}
}


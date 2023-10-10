package p15

import "os"

type spot struct {
    x, y int
}

func ImpossibleBeacon() {
	input, _ := os.Open("./pkg/p15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
}

package p17

import (
	"bufio"
	"os"
)

type coor struct {
    x, y int
}

func fallingRocks() {
    input, _ := os.Open("./pkg/p17/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    jetPattern := sc.Text()

}

func checkRockType1(start coor, tower &map[coor]bool) {
    if coor.x < 0 || coor.x > 6 {
        return tower;
    } 
}

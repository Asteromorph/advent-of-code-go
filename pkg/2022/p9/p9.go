package p9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coor struct {
	X int
	Y int
}

func TraceTail() {
	input, _ := os.Open("./pkg/p9/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	visited := make(map[Coor]bool)

	var x, y int
	tail := Coor{X: x, Y: y}
	head := Coor{X: x, Y: y}
    visited[tail] = true
	for sc.Scan() {
        dir := rune(sc.Text()[0])
        moves, _ := strconv.Atoi(sc.Text()[2:])

        for moves > 0 {
            switch dir {
            case 'R':
                head.X++
            case 'L':
                head.X--
            case 'U':
                head.Y++
            case 'D':
                head.Y--
            }
            moves--
            tail = moveTail(tail, head)
            visited[tail] = true
        }
	}
    fmt.Println(len(visited))
}

func moveTail(tail, head Coor) (newTail Coor) {
    newTail = tail
    switch(Coor{head.X - tail.X, head.Y - tail.Y}) {
    case Coor{1, 2}, Coor{2, 1}:
        newTail.X++
        newTail.Y++
    case Coor{1, -2}, Coor{2, -1}:
        newTail.X++
        newTail.Y--
    case Coor{-1, -2}, Coor{-2, -1}:
        newTail.X--
        newTail.Y--
    case Coor{-1, 2}, Coor{-2, 1}:
        newTail.X--
        newTail.Y++
    case Coor{2, 0}:
        newTail.X++
    case Coor{-2, 0}:
        newTail.X--
    case Coor{0, 2}:
        newTail.Y++
    case Coor{0, -2}:
        newTail.Y--
    }
    return
}

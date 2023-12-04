package p9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TraceTailWith10Knots() {
	input, _ := os.Open("./pkg/p9/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	visited := make(map[Coor]bool)
	knots := make([]Coor, 10)

	visited[knots[9]] = true
	for sc.Scan() {
		dir := rune(sc.Text()[0])
		moves, _ := strconv.Atoi(sc.Text()[2:])

		for moves > 0 {
			switch dir {
			case 'R':
				knots[0].X++
			case 'L':
				knots[0].X--
			case 'U':
				knots[0].Y++
			case 'D':
				knots[0].Y--
			}
                        for i := range knots[:len(knots) - 1] {
                            knots[i+1] = moveTail(knots[i+1], knots[i])
                        }
			moves--
		    fmt.Printf("%d: [%d][%d]\n",9, knots[9].X, knots[9].Y)
			visited[knots[9]] = true
		}
	}
	fmt.Println(len(visited))
}

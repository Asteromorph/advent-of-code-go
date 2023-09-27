package p8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func HighestScenicScore() {
	input, _ := os.Open("./pkg/p8/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)

	var forest [][]int
	for sc.Scan() {
		row := []int{}
		for _, height := range sc.Text() {
			h, _ := strconv.Atoi(string(height))
			row = append(row, h)
		}
		forest = append(forest, row)
	}

	rows := len(forest)
	cols := len(forest[0])
	fmt.Println(rows)
	fmt.Println(cols)
	bestView := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tmp := highestScore(i, j, forest)
			if tmp > bestView {
				bestView = tmp
			}
		}
	}
	fmt.Println(bestView)
}

func highestScore(x, y int, forest [][]int) (score int) {
	score = 1
	if x == 0 || y == 0 || x == len(forest)-1 || y == len(forest[0])-1 {
		return
	}
	fmt.Printf("curr [%d][%d]: %d\n", x, y, forest[x][y])

	for top := x - 1; top >= 0; top-- {
		if top == 0 || forest[top][y] >= forest[x][y] {
			score = score * (x - top)
			break
		}
	}

	for bottom := x + 1; bottom < len(forest); bottom++ {
		if bottom == len(forest)-1 || forest[bottom][y] >= forest[x][y] {
			score = score * (bottom - x)
			break
		}
	}

	for left := y - 1; left >= 0; left-- {
		if left == 0 || forest[x][left] >= forest[x][y] {
			score = score * (y - left)
			break
		}
	}
	for right := y + 1; right < len(forest[0]); right++ {
		if right+1 == len(forest[0]) || forest[x][right] >= forest[x][y] {
			score = score * (right - y)
			break
		}
	}

	return
}

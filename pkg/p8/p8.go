package p8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CountVisibleTrees() {
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
	visibleTree := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if checkVisibility(i, j, forest) {
				visibleTree += 1
			}
		}
	}
	fmt.Println(visibleTree)
}

func checkVisibility(x, y int, forest [][]int) (isVisible bool) {
	fmt.Printf("current tree: %d\n", forest[x][y])
	if x == 0 || y == 0 || x == len(forest)-1 || y == len(forest[0])-1 {
		return true
	}
	hidden := 0

	for top := 0; top < x; top++ {
	fmt.Printf("top tree: %d", forest[top][y])
		if forest[top][y] >= forest[x][y] {
			hidden += 1
			break
		}
	}

	for bottom := len(forest) - 1; bottom > x; bottom-- {
	fmt.Printf("bottom tree: %d", forest[bottom][y])
		if forest[bottom][y] >= forest[x][y] {
			hidden += 1
			break
		}
	}

	for left := 0; left < y; left++ {
	fmt.Printf("left tree: %d", forest[x][left])
		if forest[x][left] >= forest[x][y] {
			hidden += 1
			break
		}
	}
	for right := len(forest[0]) - 1; right > y; right-- {
	fmt.Printf("right tree: %d", forest[x][right])
		if forest[x][right] >= forest[x][y] {
			hidden += 1
			break
		}
	}
        
	fmt.Printf("hidden: %d\n", hidden)
	if hidden == 4 {
		return false
	}
	return true
}

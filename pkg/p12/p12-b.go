package p12

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x      int
	y      int
	length int
	value  rune
}

type Dir struct {
	x int
	y int
}

func getPoint() Point {
	return Point{
		x:      0,
		y:      0,
		length: 1,
	}
}

func popLeft(p *[]Point) (point Point) {
	if len(*p) == 0 {
		return Point{}
	}

	point = (*p)[0]
	*p = (*p)[1:]
	return point
}

func ShortestPath() {
	input, _ := os.Open("./pkg/p12/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
	heightMap := make([][]rune, 0)

	for sc.Scan() {
		var line []rune
		for _, height := range sc.Text() {
			if height == 'S' {
				line = append(line, 'a'-1)
				continue
			}
			if height == 'E' {
				line = append(line, 'z'+1)
				continue
			}
			line = append(line, height)
		}
		heightMap = append(heightMap, line)
	}

	directions := [4]Dir{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	turn := 0

	shortest := 0
	for row := 0; row < len(heightMap); row++ {
		for col := 0; col < len(heightMap[0]); col++{
			if heightMap[row][col] == 'a' {
				visited := map[Dir]bool{}
				queue := []Point{}

				visited[Dir{row, col}] = true
				queue = append(queue, Point{row, col, 0, 'a'})
				for len(queue) > 0 {
					cur := popLeft(&queue)
					if math.Min(float64(cur.x), float64(cur.y)) < 0 || cur.x == len(heightMap) || cur.y == len(heightMap[0]) {
						fmt.Println("continue")
						continue
					}
					fmt.Printf("%v ", cur)

					if cur.value == 'z'+1 {
						fmt.Println("found")
						if cur.length > shortest {
							shortest = cur.length
						}
						break
					}

					for _, dir := range directions {
						newX := cur.x + dir.x
						newY := cur.y + dir.y
						if !visited[Dir{newX, newY}] && math.Min(float64(newX), float64(newY)) > -1 && newX < len(heightMap) && newY < len(heightMap[0]) && heightMap[newX][newY]-heightMap[cur.x][cur.y] <= 1 {
							queue = append(queue, Point{newX, newY, cur.length + 1, heightMap[newX][newY]})
							visited[Dir{newX, newY}] = true
							// fmt.Printf("/%+v\n", queue)
						}
					}
					turn++
				}
			}

		}
	}

	fmt.Println(shortest)
}

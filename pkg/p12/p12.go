package p12

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
    x int
    y int 
    length int 
}

type Dir struct {
    x int
    y int
}

func getPoint() Point{
    return Point{
        x: 0,
        y: 0,
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
    var start, end Point

    for sc.Scan() {
        var line []rune
        for i, height := range sc.Text() {
            if height == 'S' {
                start = Point{i, len(heightMap), 1}
            }
            if height == 'E' {
                end = Point{i, len(heightMap), 1}
            }
            line = append(line, height)
        }
        heightMap = append(heightMap, line)
    }
    // fmt.Printf("%+v", heightMap)
    // fmt.Println(start)
    // fmt.Println(end)
    visited := map[Point]bool{}
    queue := []Point{} 

    dirs := 
    for {
        if len(queue) == 0 {
            break
        }

    }
}

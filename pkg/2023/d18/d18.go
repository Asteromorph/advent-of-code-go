package d18

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Coor struct {
    x, y int
}

func Part1() {
    vertex, row, col:= getPool();
    vertex = append(vertex, Coor{0, 0})
    fmt.Println("vertex", row, col)

    matrix := map[int]map[int]rune{}

    for i := 0; i <= row; i++ {
        matrix[i] = map[int]rune{}
    }

    for i := 0; i <= row; i++ {
        for j := 0; j <= col; j++ {
            matrix[i][j] = '.'
        }
    }

    var org, dest Coor
    for i := 0; i < len(vertex) - 1; i++ {
        if vertex[i].x < vertex[i + 1].x || vertex[i].y < vertex[i + 1].y {
            org, dest = vertex[i], vertex[i + 1]
        } else {
            dest, org = vertex[i], vertex[i + 1]
        }
        for r := org.x; r <= dest.x; r++ {
            matrix[r][vertex[i].y] = '#'
        }
        for c := org.y; c <= dest.y; c++ {
            matrix[vertex[i].x][c] = '#'
        }
    }
    // printMatrix(matrix)
    // colVertex := mapVertexToCol(vertex[:len(vertex) - 1])
    fmt.Println(getPoolVolume(matrix, vertex[:len(vertex) - 1], row, col))
}

func printMatrix(matrix map[int]map[int]rune) {
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]) ; j++ {
            fmt.Printf("%v", string(matrix[i][j]))
        }
        fmt.Println()
    }
} 

func mapVertexToCol(vertex []Coor) (map[int][]Coor) {
    rowVertex := make(map[int][]Coor)
    for _, v := range vertex {
        rowVertex[v.y] = append(rowVertex[v.y], v)
    }
    return rowVertex
}

func getPoolVolume(matrix map[int]map[int]rune, vertex []Coor, row, col int) int {
    vol, crossEdgeCount, vertexCount := 0, 0, 0
    for r := 0; r <= row; r++ {
        for c := 0; c <= col; c++ {
            if matrix[r][c] == '#' {
                if slices.Contains(vertex, Coor{r, c}) {
                    vertexCount++
                }
                if vertexCount % 2 == 0 {
                    crossEdgeCount++
                }
                vol++
            } else {
                // fmt.Println("edge, vertex, r, c", crossEdgeCount, vertexCount, r, c)
                if crossEdgeCount % 2 == 1 {
                    vol++
                }
            }
        }
        // fmt.Println("vol, edge, vertex", vol, crossEdgeCount, vertexCount)
        crossEdgeCount, vertexCount = 0, 0
    }
    return vol
}

func getPool() ([]Coor, int, int){
    input, _ := os.Open("./pkg/2023/d18/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    // matrix := [][]rune{}
    minX, maxX, minY, maxY := 0, 0, 0, 0
    row, col := 0, 0
    
    x, y := 0, 0
    vertex := []Coor{{0, 0}}
    for sc.Scan() {
        tokens := strings.Split(sc.Text(), " ")
        dist, _ := strconv.Atoi(tokens[1]) 
        if tokens[0] == "D" {
            x += dist
        }
        if tokens[0] == "R" {
            y += dist
        }
        if tokens[0] == "L" {
            y -= dist
        }
        if tokens[0] == "U" {
            x -= dist
        }
        if x > maxX {
            maxX = x
        }
        if x < minX {
            minX = x
        }
        if y > maxY {
            maxY = y
        }
        if y < minY {
            minY = y
        }
        vertex = append(vertex, Coor{x, y})
    }
    vertex = vertex[:len(vertex) - 1]

    absMinX := int(math.Abs(float64(minX)))
    absMinY := int(math.Abs(float64(minY)))
    row = maxX + absMinX
    col = maxY + absMinY

    for i, v := range vertex {
        vertex[i] = Coor{v.x + absMinX, v.y + absMinY}
    }

    fmt.Println(vertex)
    return vertex, row, col
}


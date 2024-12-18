package day6

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
    isObstacle bool
    isVisited bool
}

type Guard struct {
    x, y int
    dir string
}

func GuardDistictPositions() {
    matrix, start := getMatrix()
    // printMatrix(matrix)
    row, col := len(matrix), len(matrix[0])

    x, y := start.x, start.y
    dir := start.dir
    count := 1
    fmt.Println(x, y, col ,row)
    for !(x < 0 || x >= row || y < 0 || y >= col) {
        fmt.Printf("(%s %d %d %d) ", dir, x, y, count)
        if dir == "up" {
            if matrix[x - 1][y].isObstacle {
                dir = "right"
            } else {
                x--
                if x < 0 {
                    break
                }
                if matrix[x][y].isVisited == false {
                    count++
                }
                matrix[x][y] = Pos{false, true}
            }
        } else if dir == "down" {
            if matrix[x + 1][y].isObstacle {
                dir = "left"
            } else {
                x++
                if x >= row {
                    break
                }
                if matrix[x][y].isVisited == false {
                    count++
                }
                matrix[x][y] = Pos{false, true}
            }
        } else if dir == "left" {
            if matrix[x][y - 1].isObstacle {
                dir = "up"
            } else {
                y--
                if y < 0 {
                    break
                }
                if matrix[x][y].isVisited == false {
                    count++
                }
                matrix[x][y] = Pos{false, true}
            }
        } else if dir == "right" {
            if matrix[x][y + 1].isObstacle {
                dir = "down"
            } else {
                y++
                if y >= col {
                    break
                }
                if matrix[x][y].isVisited == false {
                    count++
                }
                matrix[x][y] = Pos{false, true}
            }
        }
    }

    fmt.Println(count)
}

func getMatrix() (map[int]map[int]Pos, Guard) {
    input, _ := os.Open("./pkg/2024/day6/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    x := 0

    matrix := map[int]map[int]Pos{}
    start := Guard{}
    for sc.Scan() {
        newLine := map[int]Pos{}
        for i, v := range sc.Text() {
            if v == '.' {
                newLine[i] = Pos{false, false} 
            } else if v == '#' {
                newLine[i] = Pos{true, false} 
            } else if v == '^' {
                newLine[i] = Pos{false, true} 
                start = Guard{x, i, "up"}
            }
        }
        matrix[x] = newLine
        x++
    }
    return matrix, start
}

func printMatrix(matrix map[int]map[int]Pos) {
    for _, row := range matrix {
        fmt.Println(row)
    } 
}

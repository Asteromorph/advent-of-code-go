package day6

import "fmt"

type Obs struct {
    x, y int
}

func LoopObstructionCount() {
    matrix, start := getMatrix()

    row, col := len(matrix), len(matrix[0])

    x, y := start.x, start.y
    dir := start.dir
    count := 0
    rectangle := []Obs{}
    fmt.Println(x, y, col ,row)
    for !(x < 0 || x >= row || y < 0 || y >= col) {
        fmt.Printf("(%s %d %d %d) ", dir, x, y, count)
        if dir == "up" {
            if matrix[x - 1][y].isObstacle {
                dir = "right"
                rectangle = append(rectangle, Obs{x - 1, y})
            } else {
                x--
                if x < 0 {
                    break
                }
            }
        } else if dir == "down" {
            if matrix[x + 1][y].isObstacle {
                dir = "left"
                rectangle = append(rectangle, Obs{x + 1, y})
            } else {
                x++
                if x >= row {
                    break
                }
            }
        } else if dir == "left" {
            if matrix[x][y - 1].isObstacle {
                dir = "up"
                rectangle = append(rectangle, Obs{x, y - 1})
            } else {
                y--
                if y < 0 {
                    break
                }
            }
        } else if dir == "right" {
            if matrix[x][y + 1].isObstacle {
                dir = "down"
                rectangle = append(rectangle, Obs{x, y + 1})
            } else {
                y++
                if y >= col {
                    break
                }
            }
        }

        if len(rectangle) == 3 {

        }
    }

    fmt.Println(count)
}

func isObstruct(rectangle []Obs, obsMatrix map[Obs]bool, row, col int, dir string) bool {
    v1, _, v3 := rectangle[0], rectangle[1], rectangle[2]
    var v4 Obs
    if len(rectangle) == 3 {
        if dir == "up" {
            v4.x, v4.y = v1.x - 1, v3.y
        } else if dir == "down" {
            v4.x, v4.y = v1.x + 1, v3.y
        } else if dir == "right" {
            v4.x, v4.y = v3.x, v1.y + 1
        } else if dir == "left" {
            v4.x, v4.y = v3.x, v1.y - 1
        }
    }

    if v4.x < 0 || v4.x >= row || v4.y < 0 || v4.y >= col {
        //supposed obstacle is outside of matrix
        return true
    }

    if dir == "up" {
        for y := v4.y; y < v1.y; y-- {
            if obsMatrix[Obs{v4.x, y}] {
                return true
            }
        }
        for x := v4.x; x < v3.x; x-- {
            if obsMatrix[Obs{x, v4.y}] {
                return true
            }
        }
    } else if dir == "down" {
        for y := v4.y; y > v1.y; y++ {
            if obsMatrix[Obs{v4.x, y}] {
                return true
            }
        }
        for x := v4.x; x > v3.x; x++ {
            if obsMatrix[Obs{x, v4.y}] {
                return true
            }
        }
    } else if dir == "right" {
        for y := v4.y; y > v3.y; y-- {
            if obsMatrix[Obs{v4.x, y}] {
                return true
            }
        }
        for x := v4.x; x < v1.x; x++ {
            if obsMatrix[Obs{x, v4.y}] {
                return true
            }
        }
    } else if dir == "left" {
        for y := v4.y; y > v3.y; y-- {
            if obsMatrix[Obs{v4.x, y}] {
                return true
            }
        }
        for x := v4.x; x < v1.x; x++ {
            if obsMatrix[Obs{x, v4.y}] {
                return true
            }
        }
    }


    return false
}

func getObstaclesList(matrix map[int]map[int]Pos) (res map[Obs]bool) {
    for x := 0; x < len(matrix); x++ {
        for y := 0; y < len(matrix[0]); y++ {
            if matrix[x][y].isObstacle {
                res[Obs{x, y}] = true
            }
        }
    }
    return
}


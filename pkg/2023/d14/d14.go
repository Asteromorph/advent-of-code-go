package d14

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Pos struct {
    x,y int
}

func tiltNorth(matrix map[int]map[int]rune) {
    for x := 0; x < len(matrix); x++ {
        for y := 0; y < len(matrix[0]); y++ {
            if matrix[x][y] == 'O' {
                rockSlideToNorth(x, y, matrix)
            }
        }
    }
}

func rockSlideToNorth(row, col int, matrix map[int]map[int]rune) {
    matrix[row][col] = '.'
    for row > 0 {
        if matrix[row - 1][col] != '.'{
            break
        }    
        row--
    }
    matrix[row][col] = 'O'
}

func tiltSouth(matrix map[int]map[int]rune) {
    for x := len(matrix) - 1; x >= 0; x-- {
        for y := 0; y < len(matrix[0]); y++ {
            if matrix[x][y] == 'O' {
                rockSlideToSouth(x, y, matrix)
            }
        }
    }
}

func rockSlideToSouth(row, col int, matrix map[int]map[int]rune) {
    matrix[row][col] = '.'
    for row < len(matrix) - 1 {
        if matrix[row + 1][col] != '.'{
            break
        }    
        row++
    }
    matrix[row][col] = 'O'
}

func tiltWest(matrix map[int]map[int]rune) {
    for y := 0; y < len(matrix[0]); y++ {
        for x := 0; x < len(matrix); x++ {
            if matrix[x][y] == 'O' {
                rockSlideToWest(x, y, matrix)
            }
        }
    }
}

func rockSlideToWest(row, col int, matrix map[int]map[int]rune) {
    matrix[row][col] = '.'
    for col > 0 {
        if matrix[row][col - 1] != '.'{
            break
        }    
        col--
    }
    matrix[row][col] = 'O'
}

func tiltEast(matrix map[int]map[int]rune) {
    for y := len(matrix[0]) - 1; y >= 0; y-- {
        for x := 0; x < len(matrix); x++ {
            if matrix[x][y] == 'O' {
                rockSlideToEast(x, y, matrix)
            }
        }
    }
}

func rockSlideToEast(row, col int, matrix map[int]map[int]rune) {
    matrix[row][col] = '.'
    for col < len(matrix[0]) - 1 {
        if matrix[row][col + 1] != '.'{
            break
        }    
        col++
    }
    matrix[row][col] = 'O'
}

func spin(matrix map[int]map[int]rune, cacheMap map[string]bool) bool {
    tiltNorth(matrix)
    tiltWest(matrix)
    tiltSouth(matrix)
    tiltEast(matrix)
    return checkExist(matrix, cacheMap)
}

func checkExist(matrix map[int]map[int]rune, cacheMap map[string]bool) bool {
    // println("-----------")
    // printMatrix(matrix)
    bytes, err := json.Marshal(matrix)
    if err != nil {
        log.Fatal("Error Marshalling matrix", matrix)
    } 
    key := string(bytes)

    if _, ok := cacheMap[key]; ok {
        return true
    }
    cacheMap[key] = true
    return false
}

func calculate(rounds *[]Pos, height int) (res int) {
    for _, p := range *rounds {
        res += height + 1 - p.x
    } 
    return res
}

func printMatrix(matrix map[int]map[int]rune) {
    for x := 0; x < len(matrix); x++ {
        for y := 0; y < len(matrix[0]); y++ {
            fmt.Printf("%s", string(matrix[x][y]))
        }
        fmt.Println()
    }
}

func totalLoad(matrix map[int]map[int]rune) (sum int) {
    // printMatrix(matrix)
    for x := 0; x < len(matrix); x++ {
        for y := 0; y < len(matrix[0]); y++ {
            if matrix[x][y] == 'O' {
                sum += len(matrix[0]) - x 
            }
        }
    }
    return sum
}

func Part2() {
    input, _ := os.Open("./pkg/2023/d14/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    cacheMap := make(map[string]bool)
    matrix := make(map[int]map[int]rune)
    x := 0
    for sc.Scan() {
        for y, r := range sc.Text() {
            if _, ok := matrix[x]; !ok {
                matrix[x] = make(map[int]rune)
            }
            matrix[x][y] = r
        }
        x++
    }

    var i int
    for i = 1; i <= 11; i++ {
        if spin(matrix, cacheMap) == true {
            fmt.Println(i)
            cacheMap = map[string]bool{}
        }
        // spin(matrix, cacheMap)
    }
    fmt.Println(totalLoad(matrix))
}

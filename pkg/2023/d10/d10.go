package d10

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Step struct {
    pipe Pipe
    coor Coor
}

type Pipe struct{
    shape rune
    curDir rune
}

type Coor struct {
    row int
    col int
    nextDir rune
}

func getMatrix() [][]rune{
    input, _ := os.Open("./pkg/2023/d10/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    matrix := [][]rune{}
    for sc.Scan() {
        matrix = append(matrix, []rune(sc.Text()))
    }
    return matrix
}

func findStartingPos(matrix [][]rune) (int, int){
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == 'S' {
                return i, j
            }
        }
    }
    return -1,-1
}

func FindFurthestPosInLoop() {
    matrix := getMatrix()
    fmt.Println(matrix)
    startRow, startCol := findStartingPos(matrix)

    pipeMap := map[Pipe]Coor{}
    pipeMap[Pipe{'|', 'u'}] = Coor{-1, 0, 'u'}
    pipeMap[Pipe{'|', 'd'}] = Coor{1, 0, 'd'}
    pipeMap[Pipe{'-', 'l'}] = Coor{0, -1, 'l'}
    pipeMap[Pipe{'-', 'r'}] = Coor{0, 1, 'r'}
    pipeMap[Pipe{'L', 'l'}] = Coor{-1, 0, 'u'}
    pipeMap[Pipe{'L', 'd'}] = Coor{0, 1, 'r'}
    pipeMap[Pipe{'7', 'r'}] = Coor{1, 0, 'd'}
    pipeMap[Pipe{'7', 'u'}] = Coor{0, -1, 'l'}
    pipeMap[Pipe{'F', 'u'}] = Coor{0, 1, 'r'}
    pipeMap[Pipe{'F', 'l'}] = Coor{1, 0, 'd'}
    pipeMap[Pipe{'J', 'd'}] = Coor{0, -1, 'l'}
    pipeMap[Pipe{'J', 'r'}] = Coor{-1, 0, 'u'}

    curRow := startRow
    curCol := startCol + 1
    stepCount := 0
    curPipe := Pipe{matrix[curRow][curCol], 'r'}
    var curCoor Coor
    for matrix[curRow][curCol] != 'S' {
        stepCount++
        curCoor = pipeMap[curPipe]
        curRow += curCoor.row
        curCol += curCoor.col
        curPipe.shape = matrix[curRow][curCol]
        curPipe.curDir = curCoor.nextDir
    }
    fmt.Println(math.Ceil(float64(stepCount)/2))
}



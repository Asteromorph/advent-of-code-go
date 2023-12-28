package d10

import "fmt"

func GetTotalTilesInLoop() {
    matrix := getMatrix()
    startRow, startCol := findStartingPos(matrix)
    maxRow := len(matrix)
    maxCol := len(matrix[0])

    pipeMap := getPipeMap()
    curRow := startRow
    curCol := startCol + 1

    loopMap := map[Coor]Pipe{}
    loopMap[Coor{startRow, startCol, '.'}] = Pipe{'F', '.'}
    curPipe := Pipe{matrix[curRow][curCol], 'r'}
    var mappedCoor Coor
    // fmt.Println(curRow, curCol, matrix[curRow][curCol] != 'S')
    for matrix[curRow][curCol] != 'S' {
        loopMap[Coor{curRow, curCol, '.'}] = curPipe
        mappedCoor = pipeMap[curPipe]
        curRow += mappedCoor.row
        curCol += mappedCoor.col
        curPipe.shape = matrix[curRow][curCol]
        curPipe.curDir = mappedCoor.nextDir
    }
    // fmt.Println(maxRow, maxCol)

    count := 0
    for r := 0; r < maxRow; r++ {
        for l := 0; l < maxCol; l++ {
            if getCuts(Coor{r,l,'.'}, loopMap, maxRow, maxCol) {
                fmt.Println(r,l)
                count++
            }
        }
    }
    fmt.Println(count)
}

func getCuts(curCoor Coor, loopMap map[Coor]Pipe, maxRow, maxCol int) bool{
    if _, ok := loopMap[curCoor]; ok {
        if ok {
            return false
        }
    }
    
    up := checkUp(curCoor, loopMap)
    down := checkDown(curCoor, loopMap, maxRow)
    left := checkLeft(curCoor, loopMap)
    right := checkRight(curCoor, loopMap, maxCol)

    fmt.Println(curCoor, up, down, left, right)
    if (up && down && left && right) {
        return true
    }
    return false
}

func checkUp(curCoor Coor, loopMap map[Coor]Pipe) bool {
    count := 0
    // copyCoor := curCoor
    for curCoor.row >= 0 {
        if isHorizontalCut(loopMap[curCoor].shape) {
            count++
        } 
        curCoor.row--
    }
    // if (copyCoor.row == 3 && copyCoor.col == 14) {
    //     fmt.Printf("[up %d %d %d]", copyCoor.row, copyCoor.col, count)
    // }

    if count % 2 == 1 {
        return true
    }
    return false
}

func checkDown(curCoor Coor, loopMap map[Coor]Pipe, maxRow int) bool {
    count := 0
    // copyCoor := curCoor
    for curCoor.row < maxRow {
        if loopMap[curCoor].shape == '-' {
            count++
        }
        curCoor.row++
    }
    // fmt.Println(count)
    // if (copyCoor.row == 3 && copyCoor.col == 14) {
    //     fmt.Printf("[down %d %d %d]", copyCoor.row, copyCoor.col, count)
    // }

    if count % 2 == 1 {
        return true
    }
    return false
}

func checkLeft(curCoor Coor, loopMap map[Coor]Pipe) bool {
    count := 0
    // copyCoor := curCoor
    for curCoor.col >= 0 {
        if isVerticalCut(loopMap[curCoor].shape) {
            count++
        }
        curCoor.col--
    }
    // if (copyCoor.row == 3 && copyCoor.col == 14) {
    //     fmt.Printf("[left %d %d %d]", copyCoor.row, copyCoor.col, count)
    // }
    // fmt.Println(count)

    if count % 2 == 1 {
        return true
    }
    return false
}

func checkRight(curCoor Coor, loopMap map[Coor]Pipe, maxCol int) bool {
    count := 0
    // copyCoor := curCoor
    for curCoor.col < maxCol  {
        if isVerticalCut(loopMap[curCoor].shape) {
            count++
        }
        curCoor.col++
    }
    // if (copyCoor.row == 3 && copyCoor.col == 14) {
    //     fmt.Printf("[right %d %d %d]", copyCoor.row, copyCoor.col, count)
    // }
    // fmt.Println(count)

    if count % 2 == 1 {
        return true
    }
    return false
}

func isHorizontalCut(curRune rune) bool {
    for _, r := range []rune{'-', 'F', 'L'}{
        if r == curRune {
            return true
        }
    }
    return false
}

func isVerticalCut(curRune rune) bool {
    for _, r := range []rune{'|', 'L', 'J'}{
        if r == curRune {
            return true
        }
    }
    return false
}

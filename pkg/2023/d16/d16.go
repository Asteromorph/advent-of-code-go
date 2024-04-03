package d16

import (
	"bufio"
	"fmt"
	"os"
	// "time"
)

type Pos struct {
    obj rune
    isEnergized bool
}

type Beam struct {
    x, y int
    dir rune
}

func getInput() map[int]map[int]Pos {
    input, _ := os.Open("./pkg/2023/d16/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)
    
    matrix := map[int]map[int]Pos{}
    x := 0
    for sc.Scan() {
        runes := map[int]Pos{}
        for i, r := range sc.Text() {
            runes[i] = Pos{r, false}
        }
        matrix[x] = runes
        x++
    }
    return matrix
}

func Part1() {
    cacheMap := map[Beam]bool{}
    matrix := getInput()
    beamRight(cacheMap, matrix, Beam{0, 0, 'r'})
    // fmt.Println(matrix)
    countEnergizedTiles(matrix)
}

func Part2() {
    // cacheMap := map[Beam]bool{}
    matrix := getInput()
    maxVal := 0

    // copy := copyMatrix(matrix)
    // beamDown(cacheMap, copy, Beam{0, 3, 'd'})
    // fmt.Println(countEnergizedTiles(copy))

    for x := range matrix {
        copy1 := copyMatrix(matrix)
        beamRight(map[Beam]bool{}, copy1, Beam{x, 0, 'r'})
        if val := countEnergizedTiles(copy1); val > maxVal {
            maxVal = val
        } 
        copy2 := copyMatrix(matrix)
        beamLeft(map[Beam]bool{}, copy2, Beam{x, len(matrix[0]) - 1, 'l'})
        if val := countEnergizedTiles(copy2); val > maxVal {
            maxVal = val
        } 
    }

    for y := range matrix[0] {
        c1 := copyMatrix(matrix)
        beamDown(map[Beam]bool{}, c1, Beam{0, y, 'd'})
        if val := countEnergizedTiles(c1); val > maxVal {
            maxVal = val
        } 
        c2 := copyMatrix(matrix)
        beamUp(map[Beam]bool{}, c2, Beam{len(matrix) - 1, y, 'u'})
        if val := countEnergizedTiles(c2); val > maxVal {
            maxVal = val
        } 
    }
    fmt.Println(maxVal)
}

func checkExist(cacheMap map[Beam]bool, curBeam Beam) bool {
    if _, ok := cacheMap[curBeam]; ok {
        return true
    }
    cacheMap[curBeam] = true
    return false
}

func beamRight(cacheMap map[Beam]bool, matrix map[int]map[int]Pos, curBeam Beam) {
    if isExist := checkExist(cacheMap, curBeam); isExist {
        return
    }
    x, y := curBeam.x, curBeam.y
    var curPos Pos
    for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
        curPos = matrix[x][y]
        if !curPos.isEnergized {
            matrix[x][y] = Pos{curPos.obj, true}
        }
        if (curPos.obj == '\\') {
            beamDown(cacheMap,matrix, Beam{x + 1, y, 'd'})
            break
        }
        if (curPos.obj == '/') {
            beamUp(cacheMap,matrix, Beam{x - 1, y, 'u'})
            break
        }
        if (curPos.obj == '|') {
            beamUp(cacheMap,matrix, Beam{x - 1, y, 'u'})
            beamDown(cacheMap,matrix, Beam{x + 1, y, 'd'})
            break
        }
        y++
    }
}

func beamLeft(cacheMap map[Beam]bool, matrix map[int]map[int]Pos, curBeam Beam) {
    if isExist := checkExist(cacheMap, curBeam); isExist {
        return
    }
    x, y := curBeam.x, curBeam.y
    var curPos Pos
    for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
        curPos = matrix[x][y]
        if !curPos.isEnergized {
            matrix[x][y] = Pos{curPos.obj, true}
        }
        if (curPos.obj == '\\') {
            beamUp(cacheMap, matrix, Beam{x - 1, y, 'u'})
            break
        }
        if (curPos.obj == '/') {
            beamDown(cacheMap, matrix, Beam{x + 1, y, 'd'})
            break
        }
        if (curPos.obj == '|') {
            beamUp(cacheMap, matrix, Beam{x - 1, y, 'u'})
            beamDown(cacheMap, matrix, Beam{x + 1, y, 'd'})
            break
        }
        y--
    }
}

func beamDown(cacheMap map[Beam]bool, matrix map[int]map[int]Pos, curBeam Beam) {
    if isExist := checkExist(cacheMap, curBeam); isExist {
        return
    }
    x, y := curBeam.x, curBeam.y
    var curPos Pos
    for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
        curPos = matrix[x][y]
        if !curPos.isEnergized {
            matrix[x][y] = Pos{curPos.obj, true}
        }
        if (curPos.obj == '\\') {
            beamRight(cacheMap, matrix, Beam{x, y + 1, 'r'})
            break
        }
        if (curPos.obj == '/') {
            beamLeft(cacheMap,matrix, Beam{x, y - 1, 'l'})
            break
        }
        if (curPos.obj == '-') {
            beamRight(cacheMap,matrix, Beam{x, y + 1, 'r'})
            beamLeft(cacheMap,matrix, Beam{x, y - 1, 'l'})
            break
        }
        x++
    }
}

func beamUp(cacheMap map[Beam]bool, matrix map[int]map[int]Pos, curBeam Beam) {
    if isExist := checkExist(cacheMap, curBeam); isExist {
        return
    }
    x, y := curBeam.x, curBeam.y
    var curPos Pos
    for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
        curPos = matrix[x][y]
        if !curPos.isEnergized {
            matrix[x][y] = Pos{curPos.obj, true}
        }
        if (curPos.obj == '\\') {
            beamLeft(cacheMap,matrix, Beam{x, y - 1, 'l'})
            break
        }
        if (curPos.obj == '/') {
            beamRight(cacheMap,matrix, Beam{x, y + 1, 'r'})
            break
        }
        if (curPos.obj == '-') {
            beamRight(cacheMap,matrix, Beam{x, y + 1, 'r'})
            beamLeft(cacheMap,matrix, Beam{x, y - 1, 'l'})
            break
        }
        x--
    }
}

func countEnergizedTiles(matrix map[int]map[int]Pos) (sum int) {
    for _, row := range matrix {
        for _, val := range row {
            if val.isEnergized {
                sum++
            }
        }
    }
    fmt.Println(sum)
    return
}

func copyMatrix(matrix map[int]map[int]Pos) map[int]map[int]Pos{
    copy := map[int]map[int]Pos{}
    for i, row := range matrix {
        r := map[int]Pos{}
        for j, val := range row {
            r[j] = val
        }
        copy[i] = r
    }
    return copy
}

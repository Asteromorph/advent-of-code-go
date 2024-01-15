package d11

import (
	"bufio"
	"fmt"
	"os"
)

type Coor struct {
	row int
	col int
}

func getMatrix() ([]Coor, int, int) {
	input, _ := os.Open("./pkg/2023/d11/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

    galaxiesCoors := []Coor{}
    row := 0
    col := 0
	for sc.Scan() {
        for i, v := range []rune(sc.Text()) {
            if v == '#' {
                galaxiesCoors = append(galaxiesCoors, Coor{row, i})
            }
        }
        col = len(sc.Text())
        row++
	}
	return galaxiesCoors, row - 1, col - 1
}

func getEmptyRows(galaxiesMap []Coor, maxRow int) (map[int]bool) {
    emptyRows := map[int]bool{}
    for i := 0; i <= maxRow; i++ {
        emptyRows[i] = true
    }

    for _, v := range galaxiesMap {
   if emptyRows[v.row] {
            delete(emptyRows, v.row)
        }
    }
    return  emptyRows
}

func getEmptyCols(galaxiesMap []Coor, maxCol int) (map[int]bool) {
    emptyCols := map[int]bool{}
    for i := 0; i <= maxCol; i++ {
        emptyCols[i] = true
    }

    for _, v := range galaxiesMap {
        if emptyCols[v.col] {
            delete(emptyCols, v.col)
        }
    }
    return emptyCols
}

func ShortestPathBetweenGalaxies() {
    galaxiesCoors, maxRow, maxCol := getMatrix()
    emptyRows := getEmptyRows(galaxiesCoors, maxRow)
    emptyCols := getEmptyCols(galaxiesCoors, maxCol)
    fmt.Println(galaxiesCoors, emptyRows, emptyCols)
}

func getDistanceForMap(galaxiesMap []Coor, emptyRows, emptyCols []int) int {
    for i := 0; i < len(galaxiesMap) - 1; i++ {
        for j := i + 1; j < len(galaxiesMap); j++ {

        }
    }
}

func getDistanceForOnePair(g1, g2 Coor, emptyRows, emptyCols map[int]bool) int {
    distance := abs(g2.col - g1.col) + abs(g2.row - g1.row)
    var left, right, top, bottom int
    if g1.col < g2.col {
        left = g1.col
        right = g2.col
    } else {
        left = g2.col
        right = g1.col
    }

    if g1.row < g2.row {
        top = g1.row
        bottom = g2.row
    } else {
        top = g2.row
        bottom = g1.row
    }

    for i := left + 1; i < right; i++ {
        if emptyCols[i] {
            distance++
        }
    }
    for i := top + 1; i < bottom; i++ {
        if emptyRows[i] {
            distance++
        }
    }
    return distance
}

func abs(a int) int{
    if a < 0 {
        return -a
    }
    return a
}

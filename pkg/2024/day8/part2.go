package day8

import "fmt"

func AntinodeCount2() {
    matrix := getInput()
    antennaLists := findAllAntennaLists(matrix)
    // fmt.Println(antennaLists)
    antennas := findAllAntinodes2(antennaLists, len(matrix), len(matrix[0]))
    fmt.Println(antennas, len(antennas))
}

func findAllAntinodes2(antennaList map[rune][]Pos, row, col int) (map[Pos]bool) {
    res := map[Pos]bool{}
    for _, list := range antennaList {
        ants := findAntinodesEach2(list, row, col)
        fmt.Println("ants", ants)
        for k, _ := range ants {
            if k.x < 0 || k.y < 0 || k.x >= row || k.y >= col {
                continue
            } 
            if _, ok := res[k]; !ok {
                res[k] = true
            }
        }
    }
    return res
}

func findAntinodesEach2(coors []Pos, row, col int) (map[Pos]bool) {
    res := map[Pos]bool{}
    fmt.Println("coors", coors)
    for i := 0; i < len(coors); i++ {
        for j := i + 1; j < len(coors); j++ {
            x, y := findDistance(coors[i], coors[j])
            p1x, p1y := coors[i].x, coors[i].y
            fmt.Println(x, y, p1x, p1y)
            for p1x >= 0 && p1y >= 0 && p1x < row && p1y < col {
                p1 := Pos{p1x, p1y}
                fmt.Println("pos", p1)
                if _, ok := res[p1]; !ok {
                    res[p1] = true
                }
                p1x, p1y = p1x - x, p1y - y
            }
            
            p1x, p1y = coors[i].x, coors[i].y
            for p1x >= 0 && p1y >= 0 && p1x < row && p1y < col {
                p1 := Pos{p1x, p1y}
                if _, ok := res[p1]; !ok {
                    res[p1] = true
                }
                p1x, p1y = p1x + x, p1y + y
            }
        }
    }
    return res
}

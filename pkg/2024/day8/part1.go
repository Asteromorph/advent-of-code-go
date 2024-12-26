package day8

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
    x, y int
}

type Area struct {
    antenna rune
    isAntinode bool
}

func getInput() (map[int]map[int]Area){
    input, _ := os.Open("./pkg/2024/day8/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    result := map[int]map[int]Area{}
    count := 0
    for sc.Scan() {
        line := map[int]Area{}
        for i, v := range sc.Text() {
            line[i] = Area{v, false}
        }
        result[count] = line
        count++
    }
    return result
}

func AntinodeCount() {
    matrix := getInput()
    antennaLists := findAllAntennaLists(matrix)
    // fmt.Println(antennaLists)
    antinodes := findAllAntinodes(antennaLists, len(matrix), len(matrix[0]))
    fmt.Println(antinodes , len(antinodes ))
}

func findAllAntinodes(antennaList map[rune][]Pos, row, col int) (map[Pos]bool) {
    res := map[Pos]bool{}
    for _, list := range antennaList {
        ants := findAntinodesEach(list, row, col)
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

func findAntinodesEach(coors []Pos, row, col int) (map[Pos]bool) {
    res := map[Pos]bool{}
    fmt.Println("coors", coors)
    for i := 0; i < len(coors); i++ {
        for j := i + 1; j < len(coors); j++ {
            x, y := findDistance(coors[i], coors[j])
            p1, p2 := Pos{coors[i].x - x, coors[i].y - y}, Pos{coors[j].x + x, coors[j].y + y} 
            fmt.Println(coors[i], coors[j], x, y,p1, p2,)
            if _, ok := res[p1]; !ok && p1.x >= 0 && p1.x < row {
                res[p1] = true
            }
            if _, ok := res[p2]; !ok && p2.y >= 0 && p2.y < col {
                res[p2] = true
            }
        }
    }
    return res
}

func findDistance(a, b Pos) (x, y int) {
    return b.x - a.x, b.y - a.y
}

func findAllSameTypeAtennas(matrix map[int]map[int]Area, t rune) (res []Pos) {
    for x, row := range matrix {
        for y, val := range row {
            if t == val.antenna {
                res = append(res, Pos{x, y})
            }
        } 
    }
    return  res
}

func findAllAntennaLists(matrix map[int]map[int]Area) map[rune][]Pos {
    res := map[rune][]Pos{}
    for _, row := range matrix {
        for _, antennaType := range row {
            if _, ok := res[antennaType.antenna]; !ok && antennaType.antenna != '.' {
                res[antennaType.antenna] = findAllSameTypeAtennas(matrix, antennaType.antenna)
            }
        } 
    }
    return res
}


package p18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coor struct {
    x, y, z int
}

func SurfaceArea() {
    input, _ := os.Open("./pkg/p18/input.txt")
    defer input.Close()

    sc := bufio.NewScanner(input)
    matrix := map[Coor]bool{}
    for sc.Scan() {
        tokens := strings.Split(sc.Text(), ",")
        x, _ := strconv.Atoi(tokens[0])
        y, _ := strconv.Atoi(tokens[1])
        z, _ := strconv.Atoi(tokens[2])
        matrix[Coor{x,y,z}] = true
    }
    // fmt.Println(matrix)
    // sides := 0
    // for k:= range matrix {
    //     surfaceCalculation(matrix, k, &sides)
    // }
    // fmt.Println(sides)
    countAirPockets(matrix)
}

func surfaceCalculation(matrix map[Coor]bool, newCube Coor, sides *int) {
    joinCount := 0
    adj := Coor{newCube.x + 1, newCube.y, newCube.z}
    _, ok := matrix[adj]
    if ok {
        joinCount++
    }
    adj = Coor{newCube.x - 1, newCube.y, newCube.z}
    _, ok = matrix[adj]
    if ok {
        joinCount++
    }
    adj = Coor{newCube.x, newCube.y - 1, newCube.z}
    _, ok = matrix[adj]
    if ok {
        joinCount++
    }
    adj = Coor{newCube.x, newCube.y + 1, newCube.z}
    _, ok = matrix[adj]
    if ok {
        joinCount++
    }
    adj = Coor{newCube.x, newCube.y, newCube.z + 1}
    _, ok = matrix[adj]
    if ok {
        joinCount++
    }
    adj = Coor{newCube.x, newCube.y, newCube.z - 1}
    _, ok = matrix[adj]
    if ok {
        joinCount++
    }
    *sides = *sides + 6 - joinCount
}

func countAirPockets(matrix map[Coor]bool) {
    adjMatrix := map[Coor]int{}
    for k := range matrix {
        adj := Coor{k.x + 1, k.y, k.z}
        if _, ok := adjMatrix[adj]; ok {
            adjMatrix[adj]++
        }
        adj = Coor{k.x - 1, k.y, k.z}
        if _, ok := adjMatrix[adj]; ok {
            adjMatrix[adj]++
        }
        adj = Coor{k.x, k.y + 1, k.z}
        if _, ok := adjMatrix[adj]; ok {
            adjMatrix[adj]++
        }
        adj = Coor{k.x, k.y - 1, k.z}
        if _, ok := adjMatrix[adj]; ok {
            adjMatrix[adj]++
        }
        adj = Coor{k.x, k.y, k.z + 1}
        if _, ok := adjMatrix[adj]; ok {
            adjMatrix[adj]++
        }
        adj = Coor{k.x, k.y, k.z - 1}
        if _, ok := adjMatrix[adj]; ok {
            adjMatrix[adj]++
        }
    }
    fmt.Println(adjMatrix)

    for k, v := range adjMatrix {
        if v == 6 {
            fmt.Println(k)
        }
    }
}

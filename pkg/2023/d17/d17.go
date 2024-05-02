package d17

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	// "time"

	"golang.org/x/exp/slices"
)

type Lava struct {
    x, y int
}

type LavaRun struct {
    lava Lava
    dir, opp rune
}

type Dist struct {
    val, curDirCount *int
    dir *rune
}

func Part1() {
    input, _ := os.Open("./pkg/2023/d17/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    matrix := map[int]map[int]int{}
    row := 0
    for sc.Scan() {
        line := map[int]int{}
        for i, v := range sc.Text() {
            line[i] = int(v - '0')
        }
        matrix[row] = line
        row++
    }

    dist := minimumHeatPath(matrix, Lava{0, 0})
    for k, v := range dist {
        fmt.Printf("%v: %v, %v, %v ", k, *v.val, *v.curDirCount, string(*v.dir))
    }
    lastLava := dist[Lava{12, 12}]
    fmt.Println(lastLava, *lastLava.val, len(matrix) - 1, len(matrix[0]) - 1)
}

func isValid(x, y int, matrix map[int]map[int]int) bool {
    if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
        return true
    }
    return false
}

func minInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func removeFromArray(arr []Lava, val Lava) []Lava{
    var index int
    for i, v := range arr {
        if v == val {
            index = i
        }
    }
    return append(arr[:index], arr[index+1:]...)
}

func findOpp(adjs [4]LavaRun, curDir rune) LavaRun {
    for _, dir := range adjs {
        if dir.opp == curDir {
            return dir
        }
    }
    return LavaRun{}
}

func minimumHeatPath(matrix map[int]map[int]int, startingPoint Lava) (map[Lava]Dist) {
    dist := make(map[Lava]Dist)
    var queue []Lava
    for r := range matrix {
        for c := range matrix[r] {
            cur := Lava{r, c}
            if !reflect.DeepEqual(startingPoint, cur) {
                initialRuneVal := 'n'
                intiialDirStepCount := 0
                initialIntVal := math.MaxInt
                dist[cur] = Dist{&initialIntVal, &intiialDirStepCount, &initialRuneVal}
            }
            queue = append(queue, cur)
        }
    }
    startingPointIntValue := 0
    startingPointDirStepCount := 0
    startingPointRuneValue := 'n'
    dist[startingPoint] = Dist{&startingPointIntValue, &startingPointDirStepCount , &startingPointRuneValue}
    adjs := [4]LavaRun{{Lava{0, 1}, 'r', 'l'},{Lava{0, -1}, 'l', 'r'},{Lava{1, 0}, 'd', 'u'},{Lava{-1, 0}, 'u', 'd'}}

    var curDir rune
    var countCurDir int

    for len(queue) != 0 {
        minDist := math.MaxInt
        var curLava Lava
        printDist(dist)
        for k, v := range dist {
            if slices.Contains(queue, k) && *v.val < minDist {
                minDist = *v.val
                // nextDir = *v.dir
                curLava = k
            }
        }

        queue = removeFromArray(queue, curLava)
        curDir = *dist[curLava].dir
        countCurDir = *dist[curLava].curDirCount
        fmt.Println("smallest adj", string(curDir), countCurDir, curLava, minDist)
        
        // time.Sleep(1 * time.Second)

        // fmt.Println("queue", queue)
        var nextLava Lava
        for _, adj := range adjs{
            nextX, nextY := curLava.x + adj.lava.x, curLava.y + adj.lava.y
            if (nextX < 0 || nextY < 0 || nextX >= len(matrix) || nextY >= len(matrix[0]) || (nextX == 0 && nextY == 0)) {
                continue
            }

            if reflect.DeepEqual(findOpp(adjs, curDir), curDir) {
                continue
            }

            if adj.dir == curDir {
                // fmt.Println("choosind dir", string(adj.dir), string(curDir), countCurDir)
                if countCurDir >= 2 {
                    // delete(dist, curLava)
                    continue
                } else {
                    countCurDir++
                }
            } else {
                countCurDir = 0
                curDir = adj.dir
            }

            nextLava.x = nextX
            nextLava.y = nextY

            // time.Sleep(1 * time.Second)

            alt := *dist[curLava].val + matrix[nextLava.x][nextLava.y]
            if _, ok := dist[nextLava]; ok {
                if alt < *dist[nextLava].val {
                    *dist[nextLava].val = alt
                    *dist[nextLava].curDirCount = countCurDir
                    *dist[nextLava].dir = adj.dir
                }
            }
            // fmt.Println("nextLava minDist", nextLava, countCurDir)
        }
    }

    return dist
}

func printDist(dist map[Lava]Dist) {
    for k, v := range dist {
        if *v.val != math.MaxInt {
            fmt.Printf("(%v %v)", k, *v.val)
        } else {
            fmt.Printf("(%v inf)", k)
        }
    }
}

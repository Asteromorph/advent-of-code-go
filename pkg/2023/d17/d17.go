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

    fmt.Println(minimumHeatPath(matrix, Lava{0, 0}))
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

func minimumHeatPath(matrix map[int]map[int]int, startingPoint Lava) (map[Lava]int) {
    dist := make(map[Lava]int)
    var queue []Lava
    for r := range matrix {
        for c := range matrix[r] {
            cur := Lava{r, c}
            if !reflect.DeepEqual(startingPoint, cur) {
                dist[cur] = math.MaxInt
            }
            queue = append(queue, cur)
        }
    }
    dist[startingPoint] = 0
    adjs := [4]LavaRun{{Lava{0, 1}, 'r', 'l'},{Lava{0, -1}, 'l', 'r'},{Lava{1, 0}, 'd', 'u'},{Lava{-1, 0}, 'u', 'd'}}

    var curDir rune
    var countCurDir int

    for len(queue) != 0 {
        minDist := math.MaxInt
        var curLava Lava
        for k, v := range dist {
            if slices.Contains(queue, k) && v < minDist {
                minDist = v
                curLava = k
            }
        }
        queue = removeFromArray(queue, curLava)
        // time.Sleep(1 * time.Second)

        // fmt.Println("queue", queue)
        var nextLava Lava
        for _, adj := range adjs{
            nextX, nextY := curLava.x + adj.lava.x, curLava.y + adj.lava.y
            if (nextX < 0 || nextY < 0 || nextX >= len(matrix) || nextY >= len(matrix[0]) || (nextX == 0 && nextY == 0)) {
                continue
            }
            if adj.dir == curDir {
                if countCurDir >= 3 {
                    continue
                } else {
                    countCurDir++
                }
            } else {
                countCurDir = 0
            }
            if reflect.DeepEqual(findOpp(adjs, curDir), curDir) {
                continue
            }
            fmt.Println("cur", curLava, countCurDir, string(curDir), dist[curLava])
            nextLava.x = nextX
            nextLava.y = nextY
            curDir = adj.dir

            // time.Sleep(1 * time.Second)

            alt := dist[curLava] + matrix[nextLava.x][nextLava.y]
            if alt < dist[nextLava] {
                dist[nextLava] = alt
            }
            // fmt.Println("nextLava minDist", nextLava, countCurDir)
        }
    }

    return dist
}

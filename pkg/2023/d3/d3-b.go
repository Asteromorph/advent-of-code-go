package d3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "unicode"
)

func findGearPosition(curLine string) []int {
    // fmt.Println(curLine)
    positions := []int{}
    for i, v := range curLine {
        if v == '*' {
            positions = append(positions, i)
        }
    }
    return positions
}

func checkGearForLine(curLine string, x int) []int {
    res := []int{}
    if len(curLine) == 0 {
        return res
    }

    if x < 0 || x >= len(curLine) {
        return res
    }

    left := checkLeft(curLine, x)
    right := checkRight(curLine, x)
    if unicode.IsDigit(rune(curLine[x])) {
        cur := left + string(curLine[x]) + right
        curNum, _ := strconv.Atoi(cur)
        res = append(res, curNum)
    } else {
        if len(left) > 0 {
            leftNum, _ := strconv.Atoi(left)
            res = append(res, leftNum)
        }

        if len(right) > 0 {
            rightNum, _ := strconv.Atoi(right)
            res = append(res, rightNum)
        }
    }

    // fmt.Println(res)
    return res
}

func checkLeft(curLine string, x int) string {
    n := 1
    res := ""
    for x - n >= 0 && unicode.IsDigit(rune(curLine[x - n])) {
        res = string(curLine[x - n]) + res
        n++
    }
    return res
}

func checkRight(curLine string, x int) string {
    n := 1
    res := ""
    for x + n < len(curLine) && unicode.IsDigit(rune(curLine[x + n])) {
        res += string(curLine[x + n])
        n++
    }
    return res
}

func CheckTotalGear(curLine, aboveLine, belowLine string) int {
    // fmt.Println(curLine, aboveLine, belowLine)
    sum := 0

    pos := findGearPosition(curLine)
    // fmt.Println(pos)

    if len(pos) == 0 {
        return 0
    }

    gearNum := []int{}
    for _, v := range pos {
        // fmt.Println(v)
        // fmt.Println(curLine, aboveLine, belowLine)
        gearNum = append(gearNum, checkGearForLine(aboveLine, v)...)
        gearNum = append(gearNum, checkGearForLine(belowLine, v)...)
        gearNum = append(gearNum, checkGearForLine(curLine, v)...)

        if len(gearNum) == 2 {
            fmt.Printf("[%d %d] ",gearNum[1], gearNum[0])
            // fmt.Println(gearNum)
            sum += gearNum[0] * gearNum[1]
        }
        gearNum = []int{}
    }

    return sum
}

func CheckTotalGearAllLines() {
    input, _ := os.Open("./pkg/2023/d3/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    var aboveLine, curLine, belowLine string

    sum := 0
    sc.Scan()
    curLine = sc.Text()
    sc.Scan()
    belowLine = sc.Text()
    sum += CheckTotalGear(curLine, aboveLine, belowLine)

    for sc.Scan() {
        aboveLine = curLine
        curLine = belowLine
        belowLine = sc.Text()
        sum += CheckTotalGear(curLine, aboveLine, belowLine)
    }
    aboveLine = curLine
    curLine = belowLine
    belowLine = sc.Text()
    sum += CheckTotalGear(curLine, aboveLine, belowLine)

    fmt.Println(sum)
}

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math"
// 	"os"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	// "unicode"
// )
//
//
// func readFile(filename string) []string {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalf("failed to open file: %s", err)
// 	}
//
// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)
// 	var txtlines []string
//
// 	for scanner.Scan() {
// 		txtlines = append(txtlines, scanner.Text())
// 	}
//
// 	file.Close()
//
// 	return txtlines
// }
//
// func inputToGrid(entry []string) [][]string {
// 	grid := make([][]string, len(entry))
// 	for i, line := range entry {
// 		grid[i] = strings.Split(line, "")
// 	}
// 	return grid
// }
//
// type Coord struct {
// 	row      int
// 	colStart int
// 	colEnd   int
// }
//
// //	func (c Coord) distance(other Coord) int {
// //		return math.Sqrt2(c.row-other.row) + min(min(abs(c.colStart-other.colStart), abs(c.colStart-other.colEnd)), abs(c.colStart-((other.colStart+other.colEnd)/2)))
// //	}
// func (c Coord) distance(other Coord) float64 {
// 	rowDiff := float64(c.row - other.row)
// 	colStartDiff := float64(c.colStart - other.colStart)
// 	colEndDiff := float64(c.colStart - other.colEnd)
// 	avgColDiff := float64(c.colStart - ((other.colStart + other.colEnd) / 2))
//
// 	return math.Sqrt(math.Pow(rowDiff, 2) + math.Min(math.Min(math.Pow(colStartDiff, 2), math.Pow(colEndDiff, 2)), math.Pow(avgColDiff, 2)))
// }
//
// func abs(i int) int {
// 	return int(math.Abs(float64(i)))
// }
//
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
//
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
//
// func CheckTotalGearAllLines() {
// 	// searchable lines and searchable grid forms
// 	fileLines := readFile("./pkg/2023/d3/input.txt")
// 	// fmt.Println(fileLines)
//
// 	// find all num indices in the file
// 	re := regexp.MustCompile("[0-9]+")
// 	numIndices := []Coord{}
// 	for row, line := range fileLines {
// 		foundIndices := re.FindAllStringIndex(line, -1)
// 		for _, foundIndice := range foundIndices {
// 			numIndices = append(numIndices, Coord{row, foundIndice[0], foundIndice[1]-1})
// 		}
// 	}
// 	// fmt.Println(numIndices)
//
// 	// find all gears indices in the file
// 	re2 := regexp.MustCompile("[*]")
// 	gearIndices := []Coord{}
// 	for row, line := range fileLines {
// 		foundIndices := re2.FindAllStringIndex(line, -1)
// 		for _, foundIndice := range foundIndices {
// 			gearIndices = append(gearIndices, Coord{row, foundIndice[0], foundIndice[1]-1})
// 		}
// 	}
// 	// fmt.Println(len(gearIndices))
//
// 	realGearIndices := []Coord{}
// 	acc := 0
//     n := 0
// 	for _, gear := range gearIndices {
// 		arounds := []int{}
// 		for _, num := range numIndices {
// 			res, _ := strconv.Atoi(fileLines[num.row][num.colStart:num.colEnd+1])
// 			if gear.distance(num) < 2 {
// 				arounds = append(arounds, res)
// 			}
// 		}
// 		if len(arounds) == 2 {
//             n++
// 			// fmt.Printf("gear at %d,%d is surrounded by %d and %d", gear.row, gear.colStart, arounds[0], arounds[1])
//                 fmt.Printf("[%d %d] ",arounds[1], arounds[0])
// 			realGearIndices = append(realGearIndices, gear)
// 			acc += arounds[0] * arounds[1]
// 		}
// 	}
// 	// fmt.Println(realGearIndices)
//
// 	fmt.Println(acc, n)
// }

package d3

import (
	"fmt"
	"strconv"
	"unicode"
)

func findGearPosition(curLine string) []int {
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
        leftNum, _ := strconv.Atoi(left)
        res = append(res, leftNum)
        rightNum, _ := strconv.Atoi(right)
        res = append(res, rightNum)
    }

    return res
}

func checkLeft(curLine string, x int) string {
    n := 1
    res := ""
    for x - n >= 0 && unicode.IsDigit(rune(curLine[x - n])) {
        res += string(curLine[x - n])
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

func CheckTotalGear(curLine, aboveLine, belowLine string, maxX int) int {
    // fmt.Println(curLine)
    // fmt.Println(maxX)
    sum := 0

    pos := findGearPosition(curLine)

    if len(pos) == 0 {
        return 0
    }

    gearNum := []int{}
    for i, _ := range pos {
        gearNum = append(gearNum, checkGearForLine(aboveLine, i)...)
        gearNum = append(gearNum, checkGearForLine(belowLine, i)...)
        gearNum = append(gearNum, checkGearForLine(curLine, i)...)
    }

    if len(gearNum) == 2 {
        sum = gearNum[0] * gearNum[1]
    }

    return sum
}

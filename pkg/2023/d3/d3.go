package d3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func TotalAllLines() {
    input, _ := os.Open("./pkg/2023/d3/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    sum := 0
    var curLine, aboveLine, belowLine string
    sc.Scan()
    curLine = sc.Text()
    maxX := len(sc.Text())
    sc.Scan()
    belowLine = sc.Text()
    sum += totalInLine(curLine, aboveLine, belowLine, maxX)

    for sc.Scan() {
        aboveLine = curLine
        curLine = belowLine
        belowLine = sc.Text()
        sum += totalInLine(curLine, aboveLine, belowLine, maxX)
    }

    aboveLine = curLine
    curLine = belowLine
    belowLine = ""
    sum += totalInLine(curLine, aboveLine, belowLine, maxX)
    fmt.Println(sum)
}

func totalInLine(curLine, aboveLine, belowLine string, maxX int) int {
    fmt.Println(curLine)
    sum := 0
    n := 0
    for i:= 0; i < len(curLine) - 1; i++ {
        // fmt.Println(i, string(curLine[i]))
        var curNum string
        n = 0
        for unicode.IsDigit(rune(curLine[i + n])) {
            fmt.Println(i, n)
            curNum += string(curLine[i + n])
            // fmt.Println("is digit", i, string(curLine[i]), rune(curLine[i]))
            n++
        }
        checkEngineNum := 0
        if len(curNum) > 0 {
            // fmt.Println(curNum)
            for j := range curLine {
                checkEngineNum = checkEngineNum + checkEngine(aboveLine, i+j, maxX)
                checkEngineNum = checkEngineNum + checkEngine(belowLine, i+j, maxX)
            }
            checkEngineNum = checkEngineNum + checkEngine(curLine, i-1, maxX)
            checkEngineNum = checkEngineNum + checkEngine(aboveLine, i-1, maxX)
            checkEngineNum = checkEngineNum + checkEngine(belowLine, i-1, maxX)

            checkEngineNum = checkEngineNum + checkEngine(curLine, i + len(curNum), maxX)
            checkEngineNum = checkEngineNum + checkEngine(aboveLine, i + len(curNum), maxX)
            checkEngineNum = checkEngineNum + checkEngine(belowLine, i + len(curNum), maxX)
            //check corners
        }
        
        if (checkEngineNum >= 1) {
            // fmt.Println("engine  num", checkEngineNum)
            convInt, _ := strconv.Atoi(curNum)
            sum += convInt
        }
        curNum = ""
        i += n
    }

    return sum
}

func checkEngine(curLine string, x, maxX int) int {
    if len(curLine) == 0 {
        return 0
    }

    if x < 0 || x >= maxX {
        return 0
    }

    if (curLine[x] >= 48 && curLine[x] <= 57) || curLine[x] == 46 {
        return 0
    } 
    // fmt.Println(x, string(curLine[x]), curLine[x])
    return 1
}

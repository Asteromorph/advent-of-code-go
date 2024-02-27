package d13

import (
	"bufio"
	"fmt"
	"os"
)

func isReflection(line1, line2 string) bool {
    for i := 0; i < len(line1); i++ {
        if line1[i] != line2[i] {
            return false
        }
    }
    return true
}

func flipToVertical(horizontalPattern []string) []string {
    vertialPattern := []string{}
    if len(horizontalPattern) > 0 {
        for i := 0; i < len(horizontalPattern[0]); i++ {
            t := []byte{}
            for _, str := range horizontalPattern {
                t = append(t, str[i])
            }
            line := string(t)
            vertialPattern = append(vertialPattern, line)
        }
    }
    return vertialPattern
}

func checkToEdge(i int, pattern []string) int {
    if (isReflection(pattern[i], pattern[i + 1])) {
        k := 1
        for i - k >= 0 && i + 1 + k < len(pattern) && isReflection(pattern[i - k], pattern[i + 1 + k]) {
            k++
        }

        k--
        if (i - k == 0 || i + 1 + k == len(pattern) - 1) {
            i++
            return i
        }
    }
    return 0
}

func check(pattern []string) int {
    m := len(pattern) / 2
    d := 0
    for m - d >= 1 || m + d < len(pattern) -1 {
        if res := checkToEdge(m + d, pattern); res != 0 {
            return res
        }
        if res := checkToEdge(m - d, pattern); res != 0 {
            return res
        }
        d++
    }
    return 0
}

func Part1() {
    input, _ := os.Open("./pkg/2023/d13/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    patterns := []string{}
    var verSum, horSum int
    for sc.Scan() {
        if len(sc.Text()) != 0 {
            patterns = append(patterns, sc.Text())
        } else {
            horSum += check(patterns)
            verSum += check(flipToVertical(patterns))
            patterns = []string{}
        }
    }
    res := horSum * 100 + verSum

    fmt.Println(res, horSum, verSum)
}

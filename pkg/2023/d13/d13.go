package d13

import (
	"bufio"
	"fmt"
	"os"
)

func isReflection(line1, line2 string) bool {
    fmt.Println("line", line1, line2)
    for i := 0; i < len(line1); i++ {
        if line1[i] != line2[i] {
            return false
        }
    }
    return true
}

func flipToVertical(horizontalPattern []string) []string {
    fmt.Println(horizontalPattern)
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

func check(pattern []string) int {
    for i := 0; i < len(pattern) - 2; i++ {
        if (isReflection(pattern[i], pattern[i + 1])) {
            k := 1
            for i - k >= 0 && i + 1 + k < len(pattern) && isReflection(pattern[i - k], pattern[i + 1 + k]) {
                k++
            }

            k--
            if (i - k == 0 || i + 1 + k == len(pattern) - 1) {
                return k
            }
        }
    }
    return 0
}

func Part1() {
    input, _ := os.Open("./pkg/2023/d13/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    patterns := []string{}
    for sc.Scan() {
        if len(sc.Text()) != 0 {
            patterns = append(patterns, sc.Text())
            if hor := check(patterns); hor > 0 {
                fmt.Println(hor)
            } else {
                fmt.Println(check(flipToVertical(patterns)))
            }
        }
    }
}

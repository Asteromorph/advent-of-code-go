package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func SumMultiplications() {
    input, _ := os.Open("./pkg/2024/day3/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    sum := 0
    for sc.Scan() {
        text := sc.Text()

        fmt.Println(text)
        sum += getExpressions(text)
    }
    fmt.Println(sum)

    // text := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
}

func getExpressions(input string) int {
    pattern := `mul\(\d{1,3},\d{1,3}\)`
    r, _ := regexp.Compile(pattern)
    exps := r.FindAllString(input, -1)

    inner, _ := regexp.Compile(`\d{1,3}`)
    sum := 0
    for _, v := range exps {
        ops := inner.FindAllString(v, -1)
        o1, _ := strconv.Atoi(ops[0])
        o2, _ := strconv.Atoi(ops[1])
        sum += o1 * o2
        fmt.Println(v, o1, o2)
    }
    return sum
}

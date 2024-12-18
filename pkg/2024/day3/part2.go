package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func SumMultiplications2() {
    input, _ := os.Open("./pkg/2024/day3/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    sum := 0
    for sc.Scan() {
        s := getExpressionsWithInstruction(sc.Text())
        fmt.Println(sc.Text())
        fmt.Println(s)
        sum += s
    }
    fmt.Println(sum)
}

func getExpressionsWithInstruction(input string) int {
    pattern := `mul\(\d{1,3},\d{1,3}\)`
    r, _ := regexp.Compile(pattern)
    exps := r.FindAllString(input, -1)
    posList := r.FindAllStringSubmatchIndex(input, -1)

    res := make(map[int]int)
    inner, _ := regexp.Compile(`\d{1,3}`)
    for i, v := range exps {
        ops := inner.FindAllString(v, -1)
        o1, _ := strconv.Atoi(ops[0])
        o2, _ := strconv.Atoi(ops[1])
        val := o1 * o2
        res[posList[i][0]] = val
    }
    fmt.Println(res)

    flags := getFlags(input)

    intervals, keys := combine(flags, res)
    cur := true
    result := 0
    for _, v := range keys {
        if intervals[v] == -1 {
            cur = true
            continue
        }

        if intervals[v] == -2 {
            cur = false
            continue
        }

        if cur {
            result += intervals[v]
        }
    }
    return result
}
 
func getFlags(input string) map[int]bool {
    fmt.Println("get flags", input)
    doPattern, _ := regexp.Compile(`do\(\)`)
    dontPattern, _ := regexp.Compile(`don't\(\)`)

    flags := make(map[int]bool)

    doList := doPattern.FindAllStringSubmatchIndex(input, -1)
    dontList := dontPattern.FindAllStringSubmatchIndex(input, -1)

    for _, v := range doList {
        flags[v[0]] = true
    }

    for _, v := range dontList {
        flags[v[0]] = false
    }

    return flags
}

func combine(flags map[int]bool, exps map[int]int) (map[int]int, []int) {
    res := make(map[int]int)
    for i, v := range flags {
        if v {
            res[i] = -1
        } else {
            res[i] = -2
        }
    } 

    for i, v := range exps {
        res[i] = v
    }

    keys := []int{}
    for key := range res {
        keys = append(keys, key)
    }
    sort.Ints(keys)
    return res, keys
}

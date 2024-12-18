package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func TotalMiddlePageNumbers() {
    input, _ := os.Open("./pkg/2024/day5/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    pages := map[int][]int{}

    for sc.Scan() {
        if sc.Text() == "end" {
            break
        } else {
            tokens := strings.Split(sc.Text(), "|")
            before, _ := strconv.Atoi(tokens[0])
            after, _ := strconv.Atoi(tokens[1])
            if _, ok := pages[before]; ok {
                pages[before] = append(pages[before], after)
            } else {
                pages[before] = []int{}
                pages[before] = append(pages[before], after)
            }
        }
    }

    fmt.Println(pages)

    total := 0
    for sc.Scan() {
        updates := sc.Text()
        tokens := strings.Split(updates, ",")
        order := []int{}
        for _, v := range tokens {
            p, _ := strconv.Atoi(v)
            order = append(order, p)
        }

        fmt.Println(order)

        if checkArr(pages, order) == true {
            m := len(order) / 2
            fmt.Println("true", order, m, order[m])
            total += order[m]
        }
        fmt.Println("total", total)
    }
}

func checkArr(pages map[int][]int, order []int) bool {
    flag := true
    for i := 0; i < len(order) - 1; i++ {
        // fmt.Println(pages[order[i]], order[i + 1], slices.Contains(pages[order[i]], order[i + 1]))
        if slices.Contains(pages[order[i]], order[i + 1]) {
            continue
        } else {
            flag = false
            fmt.Println("false", order)
        }
    }
    return flag
}

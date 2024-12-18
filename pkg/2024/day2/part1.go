package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountSafeReports() {
    input, _ := os.Open("./pkg/2024/day2/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    sum := 0
    for sc.Scan() {
        arr := []int{}
        tokens := strings.Split(sc.Text(), " ")
        for _, v := range tokens {
            if level, err := strconv.Atoi(v); err == nil {
                arr = append(arr, level)
            }
        }
        if (isSafeIncrease(arr) || isSafeDecrease(arr)) {
            fmt.Println(arr)
            sum +=1
        }
    }
    fmt.Println(sum)
}

func Abs(a, b int) int {
    if a - b < 0 {
        return b - a
    }
    return a - b
}

func isSafeIncrease(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i] > arr[i - 1] {
            dif2 := Abs(arr[i], arr[i-1])
            if (dif2 < 1 || dif2 > 3) {
                fmt.Println(arr[i], arr[i - 1], dif2)
                return false
            }
        } else {
            return false
        }
    }
    return true
}

func isSafeDecrease(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[i - 1] {
            dif2 := Abs(arr[i], arr[i-1])
            if (dif2 < 1 || dif2 > 3) {
                return false
            }
        } else {
            return false
        }
    }
    return true
}

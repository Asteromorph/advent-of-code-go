package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CountSafeReports2() {
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
        flag := false

        if isUp(arr) && isSafeIncrease2(arr) {
            flag = true
            sum +=1
        }

        if isDown(arr) && isSafeDecrease2(arr) {
            flag = true
            sum +=1
        }

        if !flag {
            fmt.Println(arr)
        }
    }
    fmt.Println(sum)
}

func isUp(arr []int) bool {
    if arr[0] < arr[len(arr) - 1] {
        return true
    }
    return false
}

func isDown(arr []int) bool {
    if arr[0] > arr[len(arr) - 1] {
        return true
    }
    return false
}

func isSafeIncrease2(arr []int) bool {
    count := 0
    for i := 0; i < len(arr) - 2; i++ {
        b1 := isIncrease(arr[i], arr[i + 1])
        b2 := isIncrease(arr[i + 1], arr[i + 2])

        if b1 && b2 {
            continue
        } else {
            count++
            if count > 1 {
                return false
            }

            if b3 := isIncrease(arr[i], arr[i + 2]); b3 { 
                i++
            } else {
                if !b1 {
                    count++
                }
            } 
        }
    }
    return true 
} 

func isIncrease(n1, n2 int) bool {
    dif := Abs(n1, n2)
    // fmt.Println(n1, n2)
    if n1 < n2 && (dif >= 1 && dif <= 3) {
        return true
    }
    return false
}

func isSafeDecrease2(arr []int) bool {
    count := 0
    for i := 0; i < len(arr) - 2; i++ {
        b1 := isDecrease(arr[i], arr[i + 1])
        b2 := isDecrease(arr[i + 1], arr[i + 2])
        if b1 && b2 {
            continue
        } else {
            count++
            if count > 1 {
                return false
            }

            if b3 := isDecrease(arr[i], arr[i + 2]); b3 { 
                i++
            } else {
                if !b1 {
                    count++
                }
            }
        }
    }
    return true
}

func isDecrease(n1, n2 int) bool {
    dif := Abs(n1, n2)
    // fmt.Println(n1, n2)
    if n1 > n2 && (dif >= 1 && dif <= 3) {
        return true
    }
    return false
}


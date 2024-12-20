package day7

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

func TotalCalibration2() {
    input := getInput()
    fmt.Println(input)

    res := 0
    for k, v := range input {
        res += check2(k, v)
    }
    fmt.Println(res)
}

func check2(target int, ops []int) int {
    arr := []int{ops[0]}
    for i := 1; i < len(ops); i++ {
        newVals := []int{}
        for _, arrValue := range arr {
            newVals = append(newVals, arrValue * ops[i])
            newVals = append(newVals, arrValue + ops[i])
            newVals = append(newVals, combine(arrValue, ops[i]))
        }
        arr = newVals
    } 

    if correct := slices.Index(arr, target); correct >= 0 {
        fmt.Println(correct, target, arr)
        return arr[correct]
    }
    return 0
}

func combine(a, b int) int {
    strA := strconv.Itoa(a)
    strB := strconv.Itoa(b)
    res, _ := strconv.Atoi(strA + strB)
    return res
}


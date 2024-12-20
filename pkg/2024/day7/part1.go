package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func TotalCalibration() {
    input := getInput()
    fmt.Println(input)

    res := 0
    for k, v := range input {
        res += check(k, v)
    }
    fmt.Println(res)
}

func check(target int, ops []int) int {
    arr := []int{ops[0]}
    for i := 1; i < len(ops); i++ {
        newVals := []int{}
        for _, arrValue := range arr {
            newVals = append(newVals, arrValue * ops[i])
            newVals = append(newVals, arrValue + ops[i])
        }
        arr = newVals
    } 

    if correct := slices.Index(arr, target); correct >= 0 {
        fmt.Println(correct, target, arr)
        return arr[correct]
    }
    return 0
}

func getInput() (map[int][]int){
    input, _ := os.Open("./pkg/2024/day7/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    result := map[int][]int{}
    for sc.Scan() {
        tokens := strings.Split(sc.Text(), ":")
        res, _ := strconv.Atoi(tokens[0])
        opsStr := strings.Fields(tokens[1])
        arr := []int{}
        for _, v := range opsStr {
            if n, err := strconv.Atoi(v); err == nil {
                arr = append(arr, n)
            }
        }
        result[res] = arr
    }
    return result
}

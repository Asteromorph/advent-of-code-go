package d9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(input string) []int {
    tokens := strings.Split(input, " ")
    res := []int{}
    for _, t := range tokens {
        num, _ := strconv.Atoi(t)
        res = append(res, num)
    }
    return res
}

func oneStep(nums []int, index int) {
    for i := 0; i < index; i++ {
        nums[i] = nums[i + 1] - nums[i] 
    }
}

func checkZero(nums []int, index int) bool {
    fmt.Println(nums, index)
    for i := 0; i < index; i++ {
        if nums[i] != 0 {
            return false
        }
    }
    return true
}

func OneLine(nums []int) int {
    index := len(nums) - 1
    for !checkZero(nums, index) {
        oneStep(nums, index)
        index--
    }
    sum := 0
    for i := index; i < len(nums); i++ {
        sum += nums[i]
    }
    return sum
}

func GetExtrapolateValue() {
    input, _ := os.Open("./pkg/2023/d9/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    sum := 0
    for sc.Scan() {
        line := parseLine(sc.Text())
        sum += OneLine(line)
    }
    fmt.Println(sum)
}

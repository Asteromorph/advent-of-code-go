package d9

import (
	"bufio"
	"fmt"
	"os"
)

func oneStep2(nums []int, index int) {
    temp := nums[0]
    for i := 0; i < index; i++ {
        nums[i] = nums[i + 1] - nums[i] 
    }
    nums[index] = temp
}

func OneLine2(nums []int) int {
    index := len(nums) - 1
    for !checkZero(nums, index) {
        oneStep2(nums, index)
        index--
    }
    temp := 0
    for i := index; i < len(nums); i++ {
        temp = nums[i] - temp 
    }
    return temp
}

func GetExtrapolateValue2() {
    input, _ := os.Open("./pkg/2023/d9/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    sum := 0
    for sc.Scan() {
        line := parseLine(sc.Text())
        sum += OneLine2(line)
    }
    fmt.Println(sum)
}

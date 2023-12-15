package d8

import (
	"bufio"
	"fmt"
	"os"
)


func StepsToNavigate2() {
    input, _ := os.Open("./pkg/2023/d8/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    sc.Scan()
    steps := []rune(sc.Text())
    fmt.Println(steps)
    sc.Scan()
    network := map[string]node{}
    for sc.Scan() {
        node := parseLine(sc.Text())
        network[node.value] = node
    }
    fmt.Println(network)
    
    stepsList := []int{}
    for k := range network {
        if string(k[2]) == "Z" {
            stepsList = append(stepsList, minStepsToDest(k, steps, network))
        }
    }

    fmt.Println(LCM(stepsList[0], stepsList[1], stepsList...))
}

func minStepsToDest(start string, steps []rune, network map[string]node) int {
    stepCount := 0
    curStep := start
    for string(network[curStep].value[2]) == "Z" {
        dir := steps[stepCount % len(steps)]
        if dir == 'L' {
            fmt.Println(dir)
            fmt.Println(network[curStep])
            curStep = network[curStep].left
        } else {
            fmt.Println(dir)
            fmt.Println(network[curStep])
            curStep = network[curStep].right
        }
        stepCount++
    }
    fmt.Println(stepCount)
    return stepCount
}

func GCD(a, b int) int {
    for b != 0 {
         t := b
         b = a % b
        a = t
    }
    return a
}

func LCM(a, b int, integers ...int) int {
    result := a * b / GCD(a, b)
    for i := 0; i < len(integers); i++ {
        result = LCM(result, integers[i])
    }
    return result
}

package d8

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type node struct {
    value string
    left string
    right string
}

func parseLine(input string) node {
    fmt.Println(input)
    re := regexp.MustCompile(`[a-zA-Z]+`)
    matches := re.FindAllString(input, -1)
    fmt.Println(matches)
    return node{
        value: matches[0],
        left: matches[1],
        right: matches[2],
    }
}

func StepsToNavigate() {
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
    
    stepCount := 0
    curStep := "AAA"
    goal := "ZZZ"
    for network[curStep].value != goal {
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
}

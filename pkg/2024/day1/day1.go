package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FileOpen() *bufio.Scanner {
    input, _ := os.Open("./pkg/2024/day1/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    return sc
}

func TotalDistance() {
    sc := FileOpen()
    c1, c2 := []int{}, []int{}

    for sc.Scan() {
        row := strings.Split(sc.Text(), "   ")
        fmt.Println(row)
	n1, err := strconv.Atoi(row[0])
	if err != nil {
	    log.Fatal("n1 is not a number")
	}
	c1 = append(c1, n1)
	n2, _ := strconv.Atoi(row[1])
	if err != nil {
	    log.Fatal("n1 is not a number")
	}
	c2 = append(c2, n2)
    }

    sort.Ints(c1)
    sort.Ints(c2)
    sum := 0
    for i, _ := range c1 {
	sum += abs(c2[i] - c1[i])
    }
    fmt.Println(sum)
}


func abs(a int) int {
    if a < 0 {
	return -a
    }
    return a
}

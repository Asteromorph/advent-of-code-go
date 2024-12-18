package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SimilarityScore() {
    input, _ := os.Open("./pkg/2024/day1/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    l1, l2 := map[int]int{}, map[int]int{}
    for sc.Scan() {
        row := strings.Split(sc.Text(), "   ")
	n1, err := strconv.Atoi(row[0])
	if err != nil {
	    log.Fatal("n1 is not a number")
	}
	if _, ok := l1[n1]; ok {
	    l1[n1]++
	} else {
	    l1[n1] = 1
	}

	n2, _ := strconv.Atoi(row[1])
	if err != nil {
	    log.Fatal("n1 is not a number")
	}
	if _, ok := l2[n2]; ok {
	    l2[n2]++
	} else {
	    l2[n2] = 1
	}
    }

    sum := 0
    for k, v := range l1 {
	if _, ok := l2[k]; ok {
	    sum += v * k * l2[k]
	}
    }
    
    fmt.Println(l1, l2, sum)
}

package d12

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func memoize(memoizeFunc func(string, []int) int) func(string, []int) int {
    store := make(map[string]int)

    return func(line string, runs []int) int {
	k, _ := json.Marshal([]interface{}{line, runs})
	key := string(k)

	if v, ok := store[key]; ok {
	    return v
	}

	result := memoizeFunc(line, runs)
	store[key] = result
	return result
    }
}

func sum(runs []int) int {
    res := 0
    if len(runs) > 0 {
	for _, v := range runs {
	    res += v
	}
    }
    return res
}

func countArrangements(line string, runs []int) int {
    if len(line) == 0 {
	if len(runs) == 0 {
	    return 1
	}
	return 0
    }

    if len(runs) == 0 {
	for _, v := range line {
	    if v == '#' {
		return 0
	    }
	}
	return 1
    }

    if len(line) < sum(runs) + len(runs) - 1 {
	return 0
    }

    if line[0] == '.' {
	return countArrangements(line[1:], runs)
    }

    if line[0] == '#' {
	cur := runs[0]
	for i := 0; i < cur; i++ {
	    if line[i] == '.' {
		return 0
	    }
	}
	if line[cur] == '#' {
	    return 0
	} 

	return countArrangements(line[cur + 1:], runs[1:])
    }

    return countArrangements("#" + line[1:], runs) + countArrangements("." + line[1:], runs)
}

func parseLine(line string) (string, []int) {
    tokens := strings.Split(line, " ")
    runsStr := strings.Split(tokens[1], ",")
    runs := []int{}
    for _, v := range runsStr {
	num, _ := strconv.Atoi(v)
	runs = append(runs, num)
    }
    return tokens[0], runs
}

func TotalArrangements() {
    input, _ := os.Open("./pkg/2023/d12/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    count := memoize(countArrangements)

    res := 0
    for sc.Scan() {
	line, runs := parseLine(sc.Text())
	res += count(line + ".", runs)
    }
    fmt.Println(res)
}

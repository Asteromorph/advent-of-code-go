package d12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
    sequenece string
    groups string
}

func stringToIntSlice(groups string) []int {
    result := []int{}
    if len(groups) > 0 {
	groupStrings := strings.Split(groups, ",")
	for _, v := range(groupStrings) {
	    val, _ := strconv.Atoi(v)
	    result = append(result, val)
	}
    }
    return result
}

func sum(groups []int) int {
    res := 0
    for _, v := range(groups) {
	res += v
    }
    return res
}

func parseLine(input string) Record {
    tokens := strings.Split(input, " ")

    return Record{
	sequenece: tokens[0],
	groups: tokens[1],
    }
}

func sliceFirstNum(groups string) string {
    if len(groups) < 2 {
	return groups[1:]
    } 
    return groups[2:]
}

func dfsWithCache(cache map[Record]int, seq, groups string) int {
    fmt.Println(seq, groups, cache)
    if len(groups) <= 0 {
	return 0
    }
    if cacheValue, ok := cache[Record{seq, groups}]; ok {
	return cacheValue
    }
    seqLen := len(seq)
    groupSlice := stringToIntSlice(groups)
    first := groupSlice[0]
    if seqLen - sum(groupSlice) - len(groupSlice) + 1 < 0 {
	return 0
    }
    hasHole := strings.IndexRune(seq, '.')
    if seqLen == first {
	if hasHole > -1 {
	    return 1
	} else {
	    return 0
	}
    }
    canUse := hasHole <= -1 && seq[first] != '#'
    if seq[0] == '#' {
        nextSeq := strings.TrimLeft(seq[first + 1:], ".")
	if canUse {
	    return dfsWithCache(cache, nextSeq, sliceFirstNum(groups))
	} else {
	    return 0
	}
    }
    skip := dfsWithCache(cache, strings.TrimLeft(seq[1:], "."), groups)
    newCacheKey := Record{seq, groups}
    if !canUse {
	cache[newCacheKey] = skip
	return skip
    }
    res := skip + dfsWithCache(cache, strings.TrimLeft(seq[first + 1:], "."), sliceFirstNum(groups))
    cache[newCacheKey] = res
    return res
}


func Part1() {
    input, _ := os.Open("./pkg/2023/d12/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    res := 0
    for sc.Scan() {
	cache := make(map[Record]int)
	record := parseLine(sc.Text())

	res += dfsWithCache(cache, record.sequenece, record.groups)
    }
    fmt.Println(res)
}

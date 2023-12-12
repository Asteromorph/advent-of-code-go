package d6

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type race struct {
    time, distance int
}

func parseLine(input string) []int {
    re := regexp.MustCompile("[0-9]+")
    matches := re.FindAllString(input, -1)

    var nums []int
    for _, match := range matches {
        if num, err := strconv.Atoi(match); err == nil {
            nums = append(nums, num)
        }
    }
    return nums
}

func numberOfWayToBeatARecord(race race) int {
    res := 0
    for i := 1; i <= race.time; i++ {
        if (race.time - i) * i > race.distance {
            res++
        }
    }
    return res
}

func TotalRace() {
    input, _ := os.Open("./pkg/2023/d6/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)
    
    sc.Scan()
    times := parseLine(sc.Text())
    sc.Scan()
    distances := parseLine(sc.Text())
    var races []race
    for i := range times {
        races = append(races, race{time: times[i], distance: distances[i]})
    }

    res := 1
    for _,v := range races {
        res *= numberOfWayToBeatARecord(v)
    }
    fmt.Println(res)
}

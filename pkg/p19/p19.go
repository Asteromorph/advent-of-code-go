package p19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
    ore, clay, obsidian, geode int
}

type Robot struct {
    config Config
    amount int
}

func GetAllQualityLevel() {
    input, _ := os.Open("./pkg/p19/input.txt")
    defer input.Close()

    sc := bufio.NewScanner(input) 
    for sc.Scan() {
        oreRobot := Robot{}
        clayRobot := Robot{}
        obsiRobot := Robot{}
        geodeRobot := Robot{}
        robots := []Robot{}
        var no int
        tokens := strings.Split(sc.Text(), ".")
        fmt.Sscanf(tokens[0], "Blueprint %d: Each ore robot costs %d ore", &no, &oreRobot.config.ore)
        fmt.Sscanf(tokens[1], " Each clay robot costs %d ore", &clayRobot.config.clay)
        fmt.Sscanf(tokens[2], " Each obsidian robot costs %d ore and %d clay", &obsiRobot.config.ore, &obsiRobot.config.clay)
        fmt.Sscanf(tokens[3], " Each geode robot costs %d ore and %d obsidian", &geodeRobot.config.ore, &geodeRobot.config.obsidian)
        robots = append(robots, oreRobot, clayRobot, obsiRobot, geodeRobot)
        fmt.Print(robots)
    }
}

func getQualityLevel(blueprints []Robot) int {
    // var ore, clay, obs, geode int
    for i:= 0; i < 24; i++ {
    }
    return 0
}

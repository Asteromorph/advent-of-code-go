package d5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type seedMap struct {
    dest int
    source int
    rangeMap int
}

func FarmingCorrespondance() {
    input, _ := os.Open("./pkg/2023/d5/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    //Get seeds
    var seeds []int
    sc.Scan()
    tokens := strings.Split(sc.Text(), ": ")
    seedsStrs := strings.Split(tokens[1], " ")
    for _, v := range seedsStrs {
        seedNum, _ := strconv.Atoi(v)
        seeds = append(seeds, seedNum)
    }
    fmt.Println(seeds)
    sc.Scan()
    sc.Scan()
    sc.Scan()

    var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTempt, temptToHumidity, humidityToLocation []seedMap
    var almanacMap [][]seedMap
    for len(sc.Text()) != 0 {
        seedToSoil = append(seedToSoil, parseLine(sc.Text()))
        sc.Scan()
    }
    sc.Scan()
    sc.Scan()
    for len(sc.Text()) != 0 {
        soilToFertilizer = append(soilToFertilizer, parseLine(sc.Text()))
        sc.Scan()
    }
    sc.Scan()
    sc.Scan()
    for len(sc.Text()) != 0 {
        fertilizerToWater = append(fertilizerToWater, parseLine(sc.Text()))
        sc.Scan()
    }
    sc.Scan()
    sc.Scan()
    for len(sc.Text()) != 0 {
        waterToLight = append(waterToLight, parseLine(sc.Text()))
        sc.Scan()
    }
    sc.Scan()
    sc.Scan()
    for len(sc.Text()) != 0 {
        lightToTempt = append(lightToTempt, parseLine(sc.Text()))
        sc.Scan()
    }
    sc.Scan()
    sc.Scan()
    for len(sc.Text()) != 0 {
        temptToHumidity = append(temptToHumidity, parseLine(sc.Text()))
        sc.Scan()
    }
    sc.Scan()
    sc.Scan()
    for len(sc.Text()) != 0 {
        humidityToLocation = append(humidityToLocation, parseLine(sc.Text()))
        sc.Scan()
    }
    almanacMap = append(almanacMap, seedToSoil)
    almanacMap = append(almanacMap, soilToFertilizer)
    almanacMap = append(almanacMap, fertilizerToWater)
    almanacMap = append(almanacMap, waterToLight)
    almanacMap = append(almanacMap, lightToTempt)
    almanacMap = append(almanacMap, temptToHumidity)
    almanacMap = append(almanacMap, humidityToLocation)
    // fmt.Println(almanacMap)

    shortestPath := math.MaxInt
    for _, s := range seeds {
        found := findAllMap(s, almanacMap)
        if found < shortestPath {
            shortestPath = found
        }
    }
    fmt.Println(shortestPath)
}

func findAllMap(seed int, allMaps [][]seedMap) int {
    for _, m := range allMaps {
        seed = find(seed, m)
    }
    fmt.Println(seed)
    return seed
}

func find(seed int, singleMap []seedMap) int {
    for _, v := range singleMap {
        if seed >= v.source && seed < v.source + v.rangeMap {
            // fmt.Println(seed, v.source, v.rangeMap,seed - v.source + v.dest)
            return seed - v.source + v.dest
        }
    }
    return seed
}

func parseLine(input string) seedMap {
    var dest, source, rangeMap int
    fmt.Sscanf(input,"%d %d %d", &dest, &source, &rangeMap)
    
    return seedMap{
        dest: dest,
        source: source,
        rangeMap: rangeMap,
    }
}

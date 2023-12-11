package d5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type seeds struct {
    source int
    r int
}

func FarmingCorrespondance2() {
    input, _ := os.Open("./pkg/2023/d5/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    //Get seeds
    var seedsList []seeds
    sc.Scan()
    tokens := strings.Split(sc.Text(), ": ")
    seedsStrs := strings.Split(tokens[1], " ")

    for i := 0; i < len(seedsStrs); i++ {
        fmt.Println(i)
        seedNum, _ := strconv.Atoi(seedsStrs[i])
        i++
        seedRange, _ := strconv.Atoi(seedsStrs[i])
        seedsList = append(seedsList, seeds{source: seedNum, r: seedRange})
    }
    fmt.Println(seedsList)
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
    for _, s := range seedsList {
        for i := s.source; i < s.r + s.source; i++ {
            found := findAllMap(i, almanacMap)
            if found < shortestPath {
                shortestPath = found
            }
        }
    }
    fmt.Println(shortestPath)
}


package d2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TotalPossibleGames() {
    input, _ := os.Open("./pkg/2023/d2/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    total := 0
    for sc.Scan() {
        gameNumber, isValid := parseLine(sc.Text(), 12,14,13)
        if isValid{
            total += gameNumber 
        }
    }
    fmt.Println(total)
}

func parseLine(str string, maxRed, maxBlue, maxGreen int) (int, bool) {
    gameNumber := 0

    tokens := strings.Split(str, ": ");
    fmt.Sscanf(tokens[0], "Game %d", &gameNumber)
    // fmt.Println(tokens[1])

    games := strings.Split(tokens[1], "; ")
    for _, g := range games {
        // fmt.Println(g)
        pulls := strings.Split(g, ", ")
        amount := 0
        colour := ""
        for _, p := range pulls {
            // fmt.Println(p)
            fmt.Sscanf(p, "%d %s", &amount, &colour)
            // fmt.Println(amount, colour)
            if !checkValid(maxRed, maxBlue, maxGreen, amount, colour) {
                fmt.Println(gameNumber)
                return gameNumber, false
            }
        }
    }

    return gameNumber, true
}

func checkValid(maxRed, maxBlue, maxGreen int, amount int, colour string) bool{
    res := !(colour == "red" && amount > maxRed || colour == "blue" && amount > maxBlue || colour == "green" && amount > maxGreen)
    fmt.Println(res, amount, colour)
    return res
}

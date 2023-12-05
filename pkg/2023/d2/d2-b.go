package d2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TotalOfMinimumRequirementGames() {
    input, _ := os.Open("./pkg/2023/d2/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    total := 0
    for sc.Scan() {
        r, b, g := parseLine2(sc.Text())
        total += r * b * g
    }
    fmt.Println(total)
}

func parseLine2(str string) (int, int ,int) {
    gameNumber := 0

    tokens := strings.Split(str, ": ");
    fmt.Sscanf(tokens[0], "Game %d", &gameNumber)
    // fmt.Println(tokens[1])

    games := strings.Split(tokens[1], "; ")
    var red, blue, green int
    for _, g := range games {
        // fmt.Println(g)
        pulls := strings.Split(g, ", ")
        amount := 0
        colour := ""
        for _, p := range pulls {
            // fmt.Println(p)
            fmt.Sscanf(p, "%d %s", &amount, &colour)
            if colour == "red" && amount > red {
                red = amount
            }
            if colour == "blue" && amount > blue {
                blue = amount
            }
            if colour == "green" && amount > green {
                green = amount
            }
        }
    }
    fmt.Println(red, blue, green)

    return red, blue, green
}

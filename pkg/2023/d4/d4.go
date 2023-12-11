package d4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
    num string 
    winningNums, handNums map[int]bool 
}

func CardGame() {
    input, _ := os.Open("./pkg/2023/d4/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    var card Card
    res := 0
    for sc.Scan() {
        card = parseLine(sc.Text())
        fmt.Println(card)
        res += cardWorth(card)
    }
    fmt.Println(res)
}

func parseLine(input string) Card {
    var card string
    tokens := strings.Split(input, ": ")
    fmt.Sscanf(tokens[0], "Card %s", &card)
    tokens = strings.Split(tokens[1], " | ")
    winningNums := tokens[0]
    handNums := tokens[1]
    wNums, hNums := make(map[int]bool),make(map[int]bool)

    tokens = strings.Split(winningNums, " ")
    for _, str := range tokens {
        str = strings.ReplaceAll(str, " ", "")
        if num, err := strconv.Atoi(str); err == nil {
            wNums[num] = true
        }
    }

    tokens = strings.Split(handNums, " ")
    for _, str := range tokens {
        str = strings.ReplaceAll(str, " ", "")
        if num, err := strconv.Atoi(str); err == nil {
            hNums[num] = true
        }
    }

    return Card{
        num: card,
        winningNums: wNums,
        handNums: hNums,
    }
}

func cardWorth(card Card) int {
    res := 0.5

    for k := range card.handNums {
        if card.winningNums[k] {
            fmt.Println(k)
            res *= 2
        }
    }
    if res == 0.5 {
        return 0
    }
    fmt.Println(res)
    return int(res)
}

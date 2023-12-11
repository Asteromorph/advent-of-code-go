package d4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func cardWorth2(card Card) int {
    res := 0

    for k := range card.handNums {
        if card.winningNums[k] {
            res += 1
        }
    }
    return res
}

func getCardCopies(input string, stack map[int]int) {
    card := parseLine(input)
    copiesRange := cardWorth2(card)
    cardNum, _ := strconv.Atoi(card.num)

    for i:= 1; i <= copiesRange; i++ {
        stack[cardNum + i] += stack[cardNum]
    }
    fmt.Println(copiesRange)
    fmt.Println(stack)
}

func CardGame2() {
    input, _ := os.Open("./pkg/2023/d4/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    stack := make(map[int]int)
    count := 1
    for sc.Scan() {
        stack[count]++
        fmt.Println(sc.Text())
        getCardCopies(sc.Text(), stack)
        count++
    }
    fmt.Println(stack)

    sum := 0
    for _, v := range stack {
        sum += v
    }
    fmt.Println(sum)
}

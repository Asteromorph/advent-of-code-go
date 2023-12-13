package d7

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
    cards string
    bid int
}

const (
    HighCard int = iota
    OnePair
    TwoPair
    ThreeAKind
    FullHouse
    FourAKind
    FiveAKind
)

func compareTwoHands(hand1, hand2 hand, cardOrder map[string]int) bool {
    type1 := getHandType(hand1)
    type2 := getHandType(hand2)

    // fmt.Println(hand1, hand2, type1, type2)

    if type1 == type2 {
        for i := range hand1.cards {
            // fmt.Println(string(hand1.cards[i]))
            if cardOrder[string(hand1.cards[i])] == cardOrder[string(hand2.cards[i])] {
                continue
            }
            if cardOrder[string(hand1.cards[i])] > cardOrder[string(hand2.cards[i])] {
                return false
            } else {
                return true
            }
        }
    }

    if type1 > type2 {
        return false
    } else {
        return true
    }
}

func getHandType(curHand hand) int {
    // fmt.Print(curHand)
    cardMap := make(map[string]int)
    for _, v := range curHand.cards {
        cardMap[string(v)]++
    }

    if len(cardMap) == 1 {
        return FiveAKind
    }

    if len(cardMap) == 2 {
        for _, v := range cardMap {
            if v == 1 || v == 4 {
                return FourAKind
            }
        }
        return FullHouse
    }

    if len(cardMap) == 3 {
        for _, v := range cardMap {
            if v == 3 {
                return ThreeAKind
            }
        }
        return TwoPair
    }

    if len(cardMap) == 4 {
        return OnePair
    }

    return HighCard
}

func GetTotalWinning() {
    input, _ := os.Open("./pkg/2023/d7/input.txt");
    defer input.Close()
    sc := bufio.NewScanner(input)

    handsList := []hand{}
    for sc.Scan() {
        tokens := strings.Split(sc.Text(), " ")
        bid, _ := strconv.Atoi(tokens[1])
        handsList = append(handsList, hand{cards: tokens[0], bid: bid})
    }

    cardOrders := make(map[string]int)
    cardOrders["2"]=2
    cardOrders["3"]=3
    cardOrders["4"]=4
    cardOrders["5"]=5
    cardOrders["6"]=6
    cardOrders["7"]=7
    cardOrders["8"]=8
    cardOrders["9"]=9
    cardOrders["T"]=10
    cardOrders["J"]=11
    cardOrders["Q"]=12
    cardOrders["K"]=13
    cardOrders["A"]=14

    sort.Slice(handsList, func(i, j int) bool {
        return compareTwoHands(handsList[i], handsList[j], cardOrders)
    })

    fmt.Println(handsList)
    sum := 0
    for i := range handsList {
        sum += handsList[i].bid * (i+1)
    }
    fmt.Println(sum)
}

package p11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
    items []int
    opt string
    optValue int
    division int
    true int
    false int
    inspectTime int
}

func MonkeySlinger() {
	input, _ := os.Open("./pkg/p11/input.txt")
	defer input.Close()

	sc := bufio.NewScanner(input)
    monkeys := []Monkey{}

    //parse input
    for sc.Scan() {
        sc.Scan()
        tokens := strings.Split(sc.Text(), ": ")
        monkey := Monkey{}
        for _, i := range strings.Split(tokens[1], ", ") {
            item, _ := strconv.Atoi(i) 
            monkey.items = append(monkey.items, item)
        } 
        sc.Scan()
        tokens = strings.Fields(sc.Text())
        monkey.opt = tokens[4]
        if tokens[5] == "old" {
            monkey.optValue = -1
        } else {
            v, _ := strconv.Atoi(tokens[5])
            monkey.optValue = v
        }
        sc.Scan()
        div, _ := strconv.Atoi(strings.Fields(sc.Text())[3])
        monkey.division = div
        sc.Scan()
        v, _ := strconv.Atoi(strings.Fields(sc.Text())[5])
        monkey.true = v 
        sc.Scan()
        v, _ = strconv.Atoi(strings.Fields(sc.Text())[5])
        monkey.false = v
        sc.Scan()

        // fmt.Printf("%+v\n", monkey)
        monkeys = append(monkeys, monkey)
    }

    for round := 1; round <= 20; round++ {
        oneRound(&monkeys)
    }
    for _, monkey := range monkeys{
        fmt.Printf("%+v\n", monkey.inspectTime)
    }
}

func oneRound(monkeys *[]Monkey) {
    for i, monkey := range *monkeys {
        for _, item := range monkey.items {
            optVal := monkey.optValue
            if monkey.optValue == -1 {
                optVal = item
            }
            var new int
            switch monkey.opt{
            case "+":
                new = item + optVal 
            case "*":
                new = item * optVal 
            }
            result := new
            // result := int(math.Trunc(float64(new)/ float64(3)))
            // if result != 0 {
            //     fmt.Printf("%d %d %d %d\n", new, item, optVal, result)
            // }
            if result % monkey.division == 0 {
                (*monkeys)[monkey.true].items = append((*monkeys)[monkey.true].items, result)
            } else {
                (*monkeys)[monkey.false].items = append((*monkeys)[monkey.false].items, result)
            }
            (*monkeys)[i].inspectTime ++ 

        }
        (*monkeys)[i].items = []int{}

        // fmt.Printf("%+v %v\n", *&monkey.items, i)
        // for _, monkey := range *monkeys{
        //     fmt.Printf("%+v\n", monkey.items)
        // }
    }
}


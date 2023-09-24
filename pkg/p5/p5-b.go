package p5

import (
        "bufio"
        "fmt"
        "os"
)

func (s *stack) pushGroup(r []rune){
    s.crates = append(s.crates, r...)
}

func (s *stack) popGroup(amt int) []rune{
    group := s.crates[len(s.crates)- amt : len(s.crates)]
    s.crates = s.crates[:len(s.crates) - amt]
    return group
}

func StackingCrates2() {
    input, _ := os.Open("./pkg/p5/input.txt")
    defer input.Close()

    sc := bufio.NewScanner(input)

    stacks := make([]stack, 9)

    sc.Scan()
    for sc.Text() != " 1   2   3   4   5   6   7   8   9 " {
        for i, v := range sc.Text() {
            if v != ' ' && v != '[' && v != ']' {
                stacks[i/4].addToFront(v)
            }
        }
        sc.Scan()
    }

    sc.Scan()
    for sc.Scan() {
        var amt, org, dest int
        fmt.Sscanf(sc.Text(), "move %d from %d to %d", &amt, &org, &dest)
        
        stacks[dest-1].pushGroup(stacks[org - 1].popGroup(amt))
    }

    for _, s := range stacks {
        fmt.Print(string(s.popGroup(1)))
    }
}



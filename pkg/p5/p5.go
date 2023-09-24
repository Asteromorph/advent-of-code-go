package p5

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
    crates []rune
}

func (s *stack) push(value rune) {
    s.crates = append(s.crates, value)
}

func (s *stack) pop() rune{
    last := s.crates[len(s.crates) - 1]
    s.crates = s.crates[:len(s.crates)  - 1]
    return last
}

func (s *stack) addToFront(value rune) {
    s.crates = append([]rune{value}, s.crates...) 
}

func StackingCrates() {
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
        
        for i := 0; i < amt; i++ {
            stacks[dest-1].push(stacks[org - 1].pop())
        }
    }

    for _, s := range stacks {
        fmt.Print(string(s.pop()))
    }
}

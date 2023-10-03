package p13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value   int
	element []*tree
	parent  *tree
}

func GetDistressSignal() {
	//Read input file
	input, _ := os.Open("./pkg/p13/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

    index := 1
    sum := 0
    for sc.Scan(){
        tree1 := buildTree(sc.Text())
        sc.Scan()
        tree2 := buildTree(sc.Text())

        if checkOrder(tree1, tree2) == 1 {
            sum += index
        }
        index++
        sc.Scan()
    }
    fmt.Println(sum)

}

func buildTree(text string) tree {
    root := tree{-1, []*tree{}, nil}
    temp := &root

    var curr string
    for _, r := range text {
        switch r {
        case '[':
            newTree := tree{-1, []*tree{}, temp}
            temp.element = append(temp.element, &newTree)
            temp = &newTree 
        case ']':
            if len(curr) > 0 {
                number, _ := strconv.Atoi(curr)
                temp.value = number
                curr = "" 
            }
            temp = temp.parent
        case ',':
            if len(curr) > 0 {
                number, _ := strconv.Atoi(curr)
                temp.value = number
                curr = "" 
            }
            temp = temp.parent
            newTree := tree{-1, []*tree{}, temp}
            temp.element = append(temp.element, &newTree)
            temp = &newTree
        default:
            curr += string(curr)
        }
    }
    return root
}

func checkOrder(first, second tree) int{
    switch {
    case len(first.element) == 0 && len(second.element) == 0:
        if first.value > second.value {
            return -1
        } else if first.value < second.value {
            return 1
        }
        return 0

    case first.value >= 0:
        return checkOrder(tree{-1, []*tree{&first}, nil}, second)
    case second.value >= 0:
        return checkOrder(first, tree{-1, []*tree{&second}, nil})

    default:
        var i int
        for i = 0; i < len(first.element) && i < len(second.element); i++ {
            isOrdered := checkOrder(*first.element[i], *second.element[i])
            if isOrdered != 0 {
                return isOrdered
            }
        }
        if i < len(first.element) {
            return -1
        } else if i < len(second.element) {
            return 1
        }
    }
    return 0
}


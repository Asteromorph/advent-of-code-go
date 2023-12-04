package p3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func GetCommonItem() {
    input, _ := os.Open("./pkg/p3/input.txt")
    defer input.Close()

    sc := bufio.NewScanner(input)
    sum := 0

    for sc.Scan(){
        firstComp := make(map[rune]bool)
        for _, leftChar := range sc.Text()[:len(sc.Text())/2] {
            firstComp[leftChar] = true
        }

        for _, rightChar := range sc.Text()[len(sc.Text())/2:] {
            if firstComp[rightChar] {
                sum += int(unicode.ToLower(rightChar) - 96)
                if unicode.IsUpper(rightChar) {
                    sum += 26
                }
                break
            }
        }
    }
    fmt.Println(sum)
}


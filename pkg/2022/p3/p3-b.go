package p3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func GetBadge() {
    input, _ := os.Open("./pkg/p3/input.txt")
    defer input.Close()

    sc := bufio.NewScanner(input)
    sum := 0
    for sc.Scan() {
        first := getItems(sc.Text())
        sc.Scan()
        second := getItems(sc.Text())
        sc.Scan()
        third := getItems(sc.Text())

        for item := range first {
            if second[item] && third[item] {
                sum += int(unicode.ToLower(item) - 96)
                if (unicode.IsUpper(item)) {
                    sum += 26
                }
                break
            }
        }
    }
    fmt.Println(sum)
}

func getItems(list string) map[rune]bool {
    set := make(map[rune]bool) 
    for _, item := range list {
        set[item] = true
    } 
    return set;
}

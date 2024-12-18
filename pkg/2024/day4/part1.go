package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func StringAppearances() {
    input, _ := os.Open("./pkg/2024/day4/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    group := [][]byte{}

    count := 0
    for sc.Scan() {
        cur := sc.Text()
        group = append(group, []byte(cur))
        count += horizontal(cur)

        if len(group) == 4 {
            // fmt.Println("4 lines", group)
            count += vertical(group)
            count += diagonal(group)
            group = group[1:]
        }
    }

    fmt.Println(count)
}

func horizontal(str string) (sum int) {
    p1, p2 := `XMAS`, `SAMX`

    r, _ := regexp.Compile(p1)
    sum += len(r.FindAllString(str, -1))

    r, _ = regexp.Compile(p2)
    sum += len(r.FindAllString(str, -1))
    // fmt.Println("horizontal", str, sum)
    return sum
}

func vertical(matrix [][]byte) (count int) {
    for i := 0; i < len(matrix[0]); i++ {
        col := []byte{}
        for j := 0; j < len(matrix); j++ {
            col = append(col, matrix[j][i])
        }
        str := string(col)
        // fmt.Println("vertical", str, horizontal(str))
        count += horizontal(str)
    }

    return count
}

func diagonal(matrix [][]byte) (count int) {
    for i := 0; i < len(matrix); i++ {
        fmt.Println(string(matrix[i]))
    }
    fmt.Println("---------")

    for i := 0; i < len(matrix[0]); i++ {
        if i + 3 >= len(matrix[0]) {
            continue
        }
        col := []byte{} 
        for j := 0; j < 4; j++ {
            col = append(col, matrix[j][i + j])
        }
        str := string(col)
        // fmt.Println("diagonal", str, horizontal(str))

        count += horizontal(str)
        col = []byte{} 
        for j := 0; j < 4; j++ {
            col = append(col, matrix[3 - j][i + j])
        }
        str = string(col)
        // fmt.Println("diagonal", str, horizontal(str))
        count += horizontal(str)
    }
    fmt.Println(count)
    return count
}

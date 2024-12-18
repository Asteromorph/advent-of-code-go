package day4

import (
	"bufio"
	"fmt"
	"os"
)

func StringAppearances2() {
    input, _ := os.Open("./pkg/2024/day4/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    group := [][]byte{}

    count := 0
    for sc.Scan() {
        cur := sc.Text()
        group = append(group, []byte(cur))

        if len(group) == 3 {
            // fmt.Println("4 lines", group)
            count += findX(group)
            group = group[1:]
        }
    }

    fmt.Println(count)
}

func findX(matrix [][]byte) (count int) {
    for i := 0; i < len(matrix[0]) - 2; i++ {
        if matrix[1][i + 1] == 'A' && ((matrix[0][i] == 'M' && matrix[2][i + 2] == 'S') || (matrix[0][i] == 'S' && matrix[2][i + 2] == 'M')) && ((matrix[0][i + 2] == 'M' && matrix[2][i] == 'S') || (matrix[0][i + 2] == 'S' && matrix[2][i] == 'M')) {
            count++
        }
    }
    return count
}

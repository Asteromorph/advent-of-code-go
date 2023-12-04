package d1

import (
	"bufio"
	// "fmt"
	"os"
	"unicode"
)

func fromFirst(x string) (int, int) {
    for i := 0; i < len(x); i++ {
        if unicode.IsDigit(rune(x[i])) {
            return i, int(x[i]) - 48
        }
    }
    return -1, -1 
}

func fromLast(x string) (int, int) {
    for i := len(x) - 1; i >= 0; i-- {
        if unicode.IsDigit(rune(x[i])) {
            return i, int(x[i]) - 48
        }
    }
    return -1, -1
}

func Calibration() int {
    input, _ := os.Open("./pkg/2023/d1/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    res := 0

    for sc.Scan() {
        // firstNumber := fromFirst(sc.Text())
        // secondNumber := fromLast(sc.Text())
        // fmt.Println(firstNumber, secondNumber)
        // res += firstNumber * 10 + secondNumber

    }
    return res
}

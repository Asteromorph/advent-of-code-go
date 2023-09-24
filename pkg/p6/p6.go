package p6

import (
	"bufio"
	"fmt"
	"os"
)

func GetStartBufferSignal() {
    input, _ := os.Open("./pkg/p6/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)
    sc.Scan() 

    const neededLen = 14
    for i := range sc.Text() {
        buf := make(map[byte]bool)
        for j := 0; j < neededLen; j++ {
            buf[sc.Text()[i + j]] = true
        }
        if len(buf) == neededLen {
            fmt.Println(i + neededLen)
            break
        }
    }
}

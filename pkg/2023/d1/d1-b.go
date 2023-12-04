package d1

import (
	"bufio"
	"fmt"
	"os"
)

func fromFirst2(numList map[string]int, x string) (int, int) {
    pos, val := len(x) + 1, -1
    for k := range numList {
        tmpPos, tmpVal := checkStringFromFirst(k, x)
        if tmpPos < pos && tmpPos != -1{
            pos = tmpPos
            val = numList[tmpVal]
        }
    }

    if pos == len(x) + 1 {
        return -1, -1
    }

    return pos, val
}

func checkStringFromFirst(str string, x string) (int, string) {
    for i := 0; i <= len(x) - 1 - len(str); i++ {
        count := 0
        for x[i + count] == str[count] {
            if count == len(str) - 1 {
                return i, str
            }
            count++
        }  
    }
    return -1, ""
}

func fromLast2(numList map[string]int, x string) (int, int) {
    pos, val := -1, -1
    for k := range numList {
        tmpPos, tmpVal := checkStringFromLast(k, x)
        // fmt.Println("checked", tmpPos, tmpVal, pos, val)
        if tmpPos > pos && tmpPos != -1 {
            pos = tmpPos
            val = numList[tmpVal]
        }
    }
    return pos, val
}

func checkStringFromLast(str string, x string) (int, string) {
    for i := len(x) - 1; i >= len(str) - 1; i-- {
        count := 0
        // fmt.Println("checking", string(x[i - count]), string(str[len(str) - 1 - count]), i, count)
        for x[i - count] == str[len(str) - 1 - count] {
            if count == len(str) - 1 {
                return i, str 
            }
            count++
        } 
    }

    return -1, ""
}


func Calibration2() int {
    input, _ := os.Open("./pkg/2023/d1/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)
    res := 0
    numSpelled := map[string]int{}

    numSpelled["one"]=1
    numSpelled["two"]=2
    numSpelled["three"]=3
    numSpelled["four"]=4
    numSpelled["five"]=5
    numSpelled["six"]=6
    numSpelled["seven"]=7
    numSpelled["eight"]=8
    numSpelled["nine"]=9

    firstNum, secondNum := 0, 0

    for sc.Scan() {
        pos1, first1 := fromFirst(sc.Text())
        pos2, first2 := fromFirst2(numSpelled, sc.Text())
        fmt.Println("first", pos1, first1, pos2, first2)

        if pos1 == -1 {
            firstNum = first2
        } 

        if pos2 == -1 {
            firstNum = first1
        } 

        if (pos1 != -1 && pos2 != -1) {
            if pos1 < pos2 {
                firstNum = first1
            } else {
                firstNum = first2
            }
        }


        pos1, last1 := fromLast(sc.Text())
        pos2, last2 := fromLast2(numSpelled, sc.Text())
        fmt.Println("last", pos1, last1, pos2, last2)

        if pos1 == -1 {
            secondNum = last2
        }

        if pos2 == -1 {
            secondNum = last1
        }

        if pos1 != -1 && pos2 != -1 {
            if pos1 > pos2 {
                secondNum = last1
            } else {
                secondNum = last2
            }
        }

        fmt.Println(firstNum, secondNum)
        res += firstNum * 10 + secondNum

    }
    fmt.Println(res)
    return res
    // fmt.Println(fromFirst("4nineeightseven2"))
    // fmt.Println(fromLast("4nineeightseven2"))
    // fmt.Println(fromFirst2(numSpelled, "4nineeightseven2"))
    // fmt.Println(fromLast2(numSpelled, "4nineeightseven2"))
    // return -1
}

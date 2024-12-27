package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FileSystemChecksum() {
    diskList, spaceList := getInput()
    fmt.Println(diskList,spaceList)

    curDiskLen, curSpaceLen := diskList[len(diskList) - 1], spaceList[0] 
    sum := 0
    spaceIdx := 0
    idx := diskList[0]
    spaceDif := 0

    for i := 0; i < len(diskList); i++ {
        spaceDif = curSpaceLen - curDiskLen
        tmpIdx := spaceIdx
        if spaceDif > 0 {
            for j := idx; j <= idx + curDiskLen; j++ {
                sum += (len(diskList) - i) * j
            }
            curSpaceLen -= curDiskLen
        } else if spaceDif == 0 {
            for j := idx; j <= idx + curDiskLen; j++ {
                sum += (len(diskList) - i) * j
            }
            curSpaceLen -= curDiskLen
            tmpIdx = spaceIdx + 1
        }

        if tmpIdx > spaceDif {
            i++
            for j := idx; j < idx + diskList[i]; j++ {
                sum += i * j
            }
        }

    }
    fmt.Println(sum)
}

func getInput() ([]int, []int){
    input, _ := os.Open("./pkg/2024/day9/input.txt");
    defer input.Close()

    sc := bufio.NewScanner(input)

    diskList, spaceList := []int{}, []int{}
    sc.Scan()
    diskMap := sc.Text()
    for i := 0; i < len(diskMap); i++ {
        space, _ := strconv.Atoi(string(diskMap[i]))
        diskList = append(diskList, space)
        i++
        if i >= len(diskMap) {
            break
        }
        space, _ = strconv.Atoi(string(diskMap[i]))
        spaceList = append(spaceList, space)
    }
    return diskList, spaceList
}


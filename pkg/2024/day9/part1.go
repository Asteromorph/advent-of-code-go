package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FileSystemChecksum() {
    disks, spaces := getInput()
    fmt.Println(disks, spaces)

    //idx is also value
    sum, curDiskIdx, curSpaceIdx := 0, len(disks) - 1, 0
    left, right := disks[0] - 1, diskSpaceLength(disks)
    isSpace := true
    //idx of cur disk
    i := 0

    curSpace, curDiskSpace := spaces[curSpaceIdx], disks[curDiskIdx]
    totalIdx := disks[0]
    for left != right {
        if isSpace {
            for ok := true; ok; ok = curSpace > curDiskSpace {
                fmt.Printf("1: curSpaceIdx: %d, curDiskIdx: %d, curSpace: %d, curDiskSpace: %d, left: %d ,right: %d, i: %d, totalIdx: %d, sum: %d \n", curSpaceIdx, curDiskIdx, curSpace, curDiskSpace, left ,right, i, totalIdx ,sum)

                sum += addValue(totalIdx, totalIdx + curDiskSpace, curDiskIdx)
                curSpace -= curDiskSpace
                left += curDiskSpace
                right -= curDiskSpace
                totalIdx += curDiskSpace

                curDiskIdx--
                curDiskSpace = disks[curDiskIdx]
            }
            if curSpace <= curDiskSpace {
                fmt.Printf("2: curSpaceIdx: %d, curDiskIdx: %d, curSpace: %d, curDiskSpace: %d, left: %d ,right: %d, i: %d, totalIdx: %d, sum: %d \n", curSpaceIdx, curDiskIdx, curSpace, curDiskSpace, left ,right, i, totalIdx ,sum)

                sum += addValue(totalIdx, totalIdx + curSpace, curDiskIdx)
                isSpace = false
                left += curSpace
                right -= curSpace
                totalIdx += curSpace
            }
        } else {
            fmt.Printf("3: curSpaceIdx: %d, curDiskIdx: %d, curSpace: %d, curDiskSpace: %d, left: %d ,right: %d, i: %d, totalIdx: %d, sum: %d \n", curSpaceIdx, curDiskIdx, curSpace, curDiskSpace, left ,right, i, totalIdx ,sum)

            for j := 0; j < disks[i]; j++ {
                sum += totalIdx * i
                totalIdx++
            }
            i++
            left += i
            isSpace = true
            curSpaceIdx++
            curSpace = spaces[curSpaceIdx]
        }
    }

    fmt.Println(sum)
}

func addValue(from, to, val int) (res int) {
    for i := from; i < to; i++ {
        res += i * val
    }
    fmt.Println("add value", from ,to, val, res)
    return
}

func diskSpaceLength(disks []int) (sum int){
    for _, v := range disks {
        sum += v
    } 
    return 
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


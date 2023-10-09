package p14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SandUntilSafe() {
	input, _ := os.Open("./pkg/p14/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	cave := map[rock]rune{}
	var right, deep int
	for sc.Scan() {
		tokens := strings.Split(sc.Text(), " -> ")
		for i := range tokens[:len(tokens)-1] {
			from := strings.Split(tokens[i], ",")
			to := strings.Split(tokens[i+1], ",")
			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1])
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1])

			cave[rock{toX, toY}] = '#'
			cave[rock{fromX, fromY}] = '#'

			if fromX > right {
				right = fromX
			}
			if fromY > deep {
				deep = fromY
			}

			for fromX != toX || fromY != toY {
				cave[rock{fromX, fromY}] = '#'
				if fromX < toX {
					fromX++
				}
				if fromX > toX {
					fromX--
				}
				if fromY > toY {
					fromY--
				}
				if fromY < toY {
					fromY++
				}
			}
			if fromX > right {
				right = fromX
			}
			if fromY > deep {
				deep = fromY
			}
		}
	}


	falling := true
	bottom := deep + 2
	sandCount := 0
	for falling {
		newSand := rock{x: 500, y: 0}
		if cave[newSand] == 'o' {
			falling = false
			break
		}
		for {
			cave[newSand] = '"'
			
			if newSand.y + 1 == bottom {
				break
			}
			if cave[rock{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if cave[rock{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.x--
				newSand.y++
			} else if cave[rock{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.x++
				newSand.y++
			} else {
				cave[newSand] = 'o'
				sandCount++
				break
			}
		}
		fmt.Println(sandCount)
	}
	// for y := 0; y <= deep; y++ {
	//     for x := 494; x <= right; x++ {
	//         if cave[rock{x, y}] == 0{
	//             fmt.Print(".")
	//         } else {
	//             fmt.Print(string(cave[rock{x, y}]))
	//         }
	//     }
	//     fmt.Println()
	// }
	fmt.Println(sandCount)
}

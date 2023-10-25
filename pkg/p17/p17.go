package p17

import (
	"bufio"
	"fmt"
	"os"
)

type Coor struct {
    x int
    y int
}

const (
    downHit string = "wall"
    sideHit string = "tower"
    falling string = "falling"
)

func FallingRocks() {
    input, _ := os.Open("./pkg/p17/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)
    sc.Scan()

    jetPattern := sc.Text()
    fmt.Println(jetPattern)

    rock1 := []Coor{{0,0}, {1,0}, {2,0}, {3,0}}
    rock2 := []Coor{{0,0}, {1,0}, {2,0}, {1,1}, {1,-1}}
    rock3 := []Coor{{0,0}, {1,0}, {2,0}, {2,1}, {2,2}}
    rock4 := []Coor{{0,0}, {0,1}, {0,2}, {0,3}}
    rock5 := []Coor{{0,0}, {0,1}, {1,0}, {1,1}}
    rocks := [][]Coor{}
    rocks = append(rocks, rock1)
    rocks = append(rocks, rock2)
    rocks = append(rocks, rock3)
    rocks = append(rocks, rock4)
    rocks = append(rocks, rock5)

    height := -1
    spawnX := 2
    tower := make(map[Coor]bool)
    isCurrentRockFalling := true
    count := 0
    //1 step
    isFalling := false

    for i := 0; i < 2022; i++ {
	// if its + then move the rock down 1
	spawnY := height + 4;
	curRock := Coor{spawnX, spawnY}
	rockType := i % len(rocks)

	if rockType == 1 {
	    spawnY++ 
	    curRock = Coor{spawnX, spawnY}
	}
	nextPosOfRock := curRock
	fmt.Println("begin: ", curRock, " type: ", rockType)

	fmt.Println(rockType)
	for isCurrentRockFalling {
	    fmt.Println("--------")
	    //go sideways
	    switch string(jetPattern[count % len(jetPattern)]) {
	    case "<":
		nextPosOfRock.x--
	    case ">":
		nextPosOfRock.x++
	    }
	    fmt.Printf("cur: %v, next: %v, movement: %v, count: %v\n", curRock, nextPosOfRock, string(jetPattern[count % len(jetPattern)]), count)
	    count++
	    if checkRock(nextPosOfRock, rocks[rockType], tower, !isFalling) != sideHit {
		curRock = nextPosOfRock
	    }
	    fmt.Println(curRock)

	    //falling down
	    nextPosOfRock.x = curRock.x
	    nextPosOfRock.y = curRock.y-1
	    if checkRock(nextPosOfRock, rocks[rockType], tower, isFalling) == downHit {
		height = stackRock(tower, curRock, rocks[rockType], height)
		break
	    }
	    fmt.Printf("cur: %v, next: %v\n", curRock, nextPosOfRock)
	    curRock = nextPosOfRock
	    fmt.Println(curRock)
	}
	fmt.Println(count)
	fmt.Println(tower)
    }
    fmt.Printf("Height: %d\n", height)
}

//true then tower stack
func checkRock(start Coor, rock []Coor, tower map[Coor]bool, isFalling bool) string{
    // fmt.Println("check rock")
    for _, r := range rock {
	newX := start.x + r.x
	newY := start.y + r.y
	_, ok := tower[Coor{newX, newY}]
	if (!isFalling && ok) || newY == -1 {
	    // fmt.Println("down")
	    return downHit
	}
	// fmt.Printf("newX: %d, newY: %d", newX, newY)
	// fmt.Println(!isFalling && (newX < 0 || newX > 6 || ok) )
	if isFalling && (newX < 0 || newX > 6 || ok) {
	    // fmt.Println("side")
	    return sideHit
	}
    }
    return falling
}

func stackRock(tower map[Coor]bool, start Coor, rock []Coor, newHeight int) int{
    fmt.Println("stack rock")

    for _, r := range rock {
	newX := start.x + r.x
	newY := start.y + r.y
	tower[Coor{newX, newY}] = true
	if newY > newHeight {
	    newHeight = newY
	}
    }
    return newHeight
}

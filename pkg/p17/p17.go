package p17

import (
	"bufio"
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

func fallingRocks() {
    input, _ := os.Open("./pkg/p17/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)

    jetPattern := sc.Text()

    rock1 := []Coor{Coor{0,0}, Coor{1,0}, Coor{2,0}, Coor{3,0}}
    rock2 := []Coor{Coor{0,0}, Coor{1,0}, Coor{2,0}, Coor{1,1}, Coor{1,-1}}
    rock3 := []Coor{Coor{0,0}, Coor{1,0}, Coor{2,0}, Coor{2,1}, Coor{2,2}}
    rock4 := []Coor{Coor{0,0}, Coor{0,1}, Coor{0,2}, Coor{0,3}}
    rock5 := []Coor{Coor{0,0}, Coor{0,1}, Coor{1,0}, Coor{1,1}}
    rocks := [][]Coor{}
    rocks = append(rocks, rock1)
    rocks = append(rocks, rock2)
    rocks = append(rocks, rock3)
    rocks = append(rocks, rock4)
    rocks = append(rocks, rock5)

    height := 0
    spawnX := 2
    tower := make(map[Coor]bool)
    isCurrentRockFalling := true
    count := 0
    isFalling := false

    for i := 0; i < 2022; i++ {
	// if its + then move the rock down 1
	spawnY := height + 4;
	newRock := Coor{spawnX, spawnY}
	for isCurrentRockFalling {
	    rockType := count % len(rocks) 
	    if rockType == 1 {
		spawnY-- 
		newRock = Coor{spawnX, spawnY}
	    }
	    fallingType := checkRock(newRock, rocks[rockType], tower, isFalling)
	    if fallingType == sideHit {
		
	    }
	}
    }
}

//true then tower stack
func checkRock(start Coor, rock []Coor, tower map[Coor]bool, isFalling bool) string{
    for _, r := range rock {
	newX := start.x + r.x
	newY := start.y + r.y
	if isFalling && tower[Coor{newX, newY}] {
	    return downHit
	}
	if !isFalling && (newX < 0 || newY > 6 || tower[Coor{newX, newY}]) {
	    return sideHit
	}
    }
    return falling
}

func stackRock(tower map[Coor]bool, start Coor, rock []Coor, newHeight int) int{
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

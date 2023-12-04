package p15

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type zone struct {
    top, bottom, left, right int
}

func TuningFrequency() {
	input, _ := os.Open("./pkg/p15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

    spots := make(map[spot]bool)
    var left, right, top, bottom int
    for sc.Scan() {
	var sensorX, sensorY, beaconX, beaconY int
	fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

	toX := int(math.Abs(float64(sensorX) - float64(beaconX)))
	toY := int(math.Abs(float64(sensorY) - float64(beaconY)))
	toXY := toX + toY
	fmt.Println(toXY)
	
	if sensorX + toXY > right {
	    right = sensorX + toXY
	}

	if sensorX - toXY < left {
	    left = sensorX - toXY
	}

	if sensorY + toXY > bottom {
	    bottom = sensorY + toXY
	}

	if sensorY - toXY < top {
	    top = sensorY - toXY
	}

	for n:=0; n <= toXY; n++ {
	    for m := 0; m <= toXY - n; m++ {
		spots[spot{sensorX + m, sensorY + n}] = true
		spots[spot{sensorX + m, sensorY - n}] = true
		spots[spot{sensorX - m, sensorY + n}] = true
		spots[spot{sensorX - m, sensorY - n}] = true
	    }
	} 
    }
    for x := left; x <= right; x++ {
	for y:= top; y <= bottom; y++ {
	    if x < 0 || y < 0 || x > 4000000 || y > 4000000 {
		continue
	    }
	    if spots[spot{x + 1, y}] && spots[spot{x - 1, y}] && !spots[spot{x, y}] {
		fmt.Println(x, y)
	    }
	}
    }
}

package p15

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func TuningFrequency() {
	input, _ := os.Open("./pkg/p15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

    spots := make(map[spot]true)
    for sc.Scan() {
	var sensorX, sensorY, beaconX, beaconY int
	fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

	toX := int(math.Abs(float64(sensorX) - float64(beaconX)))
	toY := int(math.Abs(float64(sensorY) - float64(beaconY)))
	toXY := toX + toY

	for n:=0; n <= toXY; n++ {
	    spots[spot{sensorX, sensorY + n}] = true
	    for m := 0; m <= toXY - n; m++ {
		spots[spot{sensorX + m, toXY - }]
	    }
	} 
    }
	// fmt.Println(line10)
	// fmt.Println(toLine10)
}

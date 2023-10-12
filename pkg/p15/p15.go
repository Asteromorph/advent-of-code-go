package p15

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type spot struct {
    x, y int
}

func ImpossibleBeaconPositions() {
	input, _ := os.Open("./pkg/p15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

    line10 := make(map[spot]bool)
    targetRow := 2000000

    for sc.Scan() {
	var sensorX, sensorY, beaconX, beaconY int
	fmt.Sscanf(sc.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

	toX := int(math.Abs(float64(sensorX) - float64(beaconX)))
	toY := int(math.Abs(float64(sensorY) - float64(beaconY)))
	toTargetLine := int(math.Abs(float64(targetRow) - float64(sensorY)))

	for x:=0; x <= (toX + toY - toTargetLine); x++ {
	    line10[spot{sensorX + x, targetRow}] = true
	    line10[spot{sensorX - x, targetRow}] = true
	} 
	if beaconY == targetRow {
	    delete(line10, spot{beaconX, beaconY})
	}
    }
	fmt.Println(len(line10))
}

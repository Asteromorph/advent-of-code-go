package main

import (
	"fmt"

	"github.com/Asteromorph/advent-of-code/pkg/p1"
)

func main() {
    value, err := p1.GetMaxCalories();    
    if err != nil {
        fmt.Printf("Error %s", err)
        return;
    } 
    fmt.Printf("value: %v", value)
}

package p1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInput() (string, error) {
    data, err := os.ReadFile("input.txt")
    if err != nil {
	return "", fmt.Errorf("read error : %v", err)
    }
    return string(data), nil
}

func GetMaxCalories() (int, error) {
    input, err := GetInput();

    if err != nil {
	return 0, fmt.Errorf("%w", err)
    }
    fmt.Printf("input: %v", input)

    numberOfElves := strings.Split(input, "\n\n");
    m := 0
    for _, v := range numberOfElves {
	curr := 0
	for _, v1 := range strings.Split(v, "\n") {
	    conv, err := strconv.Atoi(v1)	     
	    if err != nil {
		fmt.Errorf("Cant convert string %v to int", v1)
	    } else {
		curr += conv
	    }
	}
	if curr > m {
	    m = curr
	} 
    }
    return m, nil
} 

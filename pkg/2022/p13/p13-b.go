package p13

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func GetDecoderKey() {
	input, _ := os.Open("./pkg/p13/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	
	var packages []tree
	for sc.Scan() {
		packages = append(packages, buildTree(sc.Text()))
		sc.Scan()
		packages = append(packages, buildTree(sc.Text()))
		sc.Scan()
	}
		packages = append(packages, buildTree("[[2]]"))
		packages = append(packages, buildTree("[[6]]"))

	sort.Slice(packages, func(i, j int) bool {
		return checkOrder(packages[i], packages[j]) == 1
	})
	key := 1
	for i, p := range packages{
		if checkOrder(p, buildTree("[[2]]")) == 0 || checkOrder(p, buildTree("[[6]]")) == 0 {
			key *= i + 1
		}
	}
	fmt.Println(key)
}

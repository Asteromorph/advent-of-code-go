
package p7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetDeleteDir() {
	input, _ := os.Open("./pkg/p7/input.txt")
	defer input.Close()
    var curObj *Object

	sc := bufio.NewScanner(input)
    fileSystem := []*Object{}

    for sc.Scan() {
        line := strings.Fields(sc.Text())
        //command line
        if len(line) > 2 {
            if line[2] == ".." {
                curObj = curObj.Parent
            } else if line[2] == "/" {
                curObj = &Object{0, "/", map[string]*Object{}, nil, false}
            } else {
                curObj = curObj.Children[line[2]]
            }
        }

        if len(line) == 2 {
            if line[0] == "dir" {
                curObj.Children[line[1]] = &Object{0, line[1], map[string]*Object{}, curObj, false}
                fileSystem = append(fileSystem, curObj.Children[line[1]])
            } else if line[0] != "$" {
                size, _ := strconv.Atoi(line[0])
                curObj.Children[line[1]] = &Object{size, line[1], map[string]*Object{}, curObj, true}
            }
        }
    }

    deficit := 30000000 - (70000000- getSize(*fileSystem[0]))
    smallestSize := getSize(*fileSystem[0])

    for _, dir := range fileSystem {
        size := getSize(*dir)
        if size > deficit && size-deficit < smallestSize-deficit {
            smallestSize = size
        }
    }
    fmt.Println(smallestSize)
}

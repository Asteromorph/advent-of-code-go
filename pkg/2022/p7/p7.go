package p7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Object struct {
    Size int
    Name string
    Children map[string]*Object
    Parent *Object
    isFile bool
}

func GetLimitDir() {
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

    total := 0
    for _, dir := range fileSystem {
        size := getSize(*dir)
        if size <= 100000 {
            total += size
        }
    }

    fmt.Println(total)

}

func getSize(root Object) (size int) {
    if root.isFile {
        return root.Size
    } else {
        for _, child := range(root.Children) {
            size += getSize(*child)
        }
    }
    return 
}

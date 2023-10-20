package p16

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type valve struct {
    name string
    flowrate int
    leadTo []string
}

func MostPressurePath() {
    graph := makeGraph()

    visited := map[string]bool{"AA": true}
    for name, valve := range graph {
	if valve.flowrate == 0 {
	    visited[name] = true
	}
    }
    fmt.Println(bfs(graph, visited, map[string]int{}, "AA", 30, 0))
}

func bfs(graph map[string]valve, visited map[string]bool, memo map[string]int, currentValve string, timeLeft, currentPressure int ) int {
    if timeLeft == 0 {
	return 0
    }

    key := hash(currentValve, timeLeft, currentPressure, visited)
    if v, ok := memo[key]; ok {
	return v
    }

    bestFlow := 0

    if !visited[currentValve] {
	visited[currentValve] = true

	newPressure := currentPressure + graph[currentValve].flowrate
	
	//Pressure if stay
	stayNewPressure := currentPressure + bfs(graph, visited, memo, currentValve, timeLeft - 1, currentPressure)

	bestFlow = int(math.Max(float64(newPressure), float64(stayNewPressure)))
	//backtrack
	visited[currentValve] = false
    }

    for _, lead := range graph[currentValve].leadTo {
	newPressure := currentPressure + bfs(graph, visited, memo, lead, timeLeft - 1, currentPressure)
	bestFlow = int(math.Max(float64(newPressure), float64(bestFlow)))
    } 
    memo[key] = bestFlow

    return bestFlow
}

func hash(currentValve string, timeLeft, currentPressure int, visited map[string]bool) string {
    valves := []string{}
    for v := range visited {
	valves = append(valves, v)
    }
    sort.Strings(valves)
    return fmt.Sprint(currentValve, timeLeft, valves, currentPressure)
}

func makeGraph() map[string]valve {
    input, _ := os.Open("./pkg/p16/input.txt")
    defer input.Close()
    sc := bufio.NewScanner(input)
    
    v := valve{}
    
    graph := map[string]valve{}

    for sc.Scan() {
	tokens := strings.Split(sc.Text(), "; ")
        _, err := fmt.Sscanf(tokens[0], "Valve %s has flow rate=%d", &v.name, &v.flowrate)    
	if err != nil {
	    panic("Parsing error")
	}
	connections := strings.Split(tokens[1], ", ")
	connections[0] = connections[0][len(connections[0])-2:]
	v.leadTo = connections;
	graph[v.name] = v
    }

    return graph
}

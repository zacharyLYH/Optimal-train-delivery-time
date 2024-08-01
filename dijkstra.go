package main

import (
	"container/heap"
	"math"
)

type Node struct {
	station string
	cost    int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type PathResult struct {
    Cost int
    Path []string
}

var shortestPathMemo = make(map[string]PathResult)

func generateKey(station1, station2 string) string {
    return station1 + "_" + station2
}


func dijkstra(start, end string, edges []Edge) PathResult {
    // Check if the result is already in the memoization map
    key := generateKey(start, end)
    if val, found := shortestPathMemo[key]; found {
        return val
    }

    // Create a map to store the shortest path to each station
    shortestPath := make(map[string]int)
    predecessor := make(map[string]string)
    for _, edge := range edges {
        shortestPath[edge.station1] = math.MaxInt64
        shortestPath[edge.station2] = math.MaxInt64
    }
    shortestPath[start] = 0

    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, Node{station: start, cost: 0})

    for pq.Len() > 0 {
        node := heap.Pop(pq).(Node)
        currentStation := node.station
        currentCost := node.cost

        // If we reached the end station, build the path, store the result in the memoization map and return the cost
        if currentStation == end {
            path := []string{}
            for at := end; at != ""; at = predecessor[at] {
                path = append([]string{at}, path...)
            }
            result := PathResult{Cost: currentCost, Path: path}
            shortestPathMemo[key] = result
            return result
        }

        // Iterate through all edges to find adjacent stations
        for _, edge := range edges {
            if edge.station1 == currentStation || edge.station2 == currentStation {
                neighbor := edge.station2
                if edge.station2 == currentStation {
                    neighbor = edge.station1
                }

                // Calculate the new cost to reach the neighbor
                newCost := currentCost + edge.timeTaken
                if newCost < shortestPath[neighbor] {
                    shortestPath[neighbor] = newCost
                    predecessor[neighbor] = currentStation
                    heap.Push(pq, Node{station: neighbor, cost: newCost})
                }
            }
        }
    }

    return PathResult{Cost: math.MaxInt64, Path: []string{}}
}

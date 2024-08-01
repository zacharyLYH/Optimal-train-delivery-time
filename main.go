package main

import (
	"fmt"
	"math"
)

type Edge struct {
	station1, station2 string
	timeTaken          int
}

type Package struct {
	name string
	weight           int
	start, end string
}

type Train struct {
	name, currentPosition string
	capacity, remaining   int
	loadedPackages        []Package
}

func copyTrain(t Train) Train {
	packagesCopy := make([]Package, len(t.loadedPackages))
	copy(packagesCopy, t.loadedPackages)

	return Train{
		name:            t.name,
		currentPosition: t.currentPosition,
		capacity:        t.capacity,
		remaining:       t.remaining,
		loadedPackages:  packagesCopy,
	}
}

// Function to calculate travel time for a train given its route
func calculateTravelTime(train Train, edges []Edge) int {
	totalTime := 0
	trainCpy := copyTrain(train)
	currentPosition := trainCpy.currentPosition
	allDestinations := make(map[string]bool)
	for _, loc := range train.loadedPackages {
		allDestinations[loc.end] = true
		allDestinations[loc.start] = true
	}

	for len(allDestinations) > 0 {
		minDistance := math.MaxInt64
		closestPkgIndex := -1

		for i, pkg := range trainCpy.loadedPackages {
			res := dijkstra(currentPosition, pkg.start, edges)
			toStart := res.Cost
			if toStart < minDistance {
				minDistance = toStart
				closestPkgIndex = i
			}
		}

		// Move to the closest package's start point
		closestPkg := trainCpy.loadedPackages[closestPkgIndex]
		res := dijkstra(currentPosition, closestPkg.start, edges)
		totalTime += res.Cost
		currentPosition = closestPkg.start
		delete(allDestinations, currentPosition)

		// Deliver all packages that can be delivered along the way
		firstDestination := dijkstra(currentPosition, trainCpy.loadedPackages[0].end, edges)
		for _, station := range firstDestination.Path[1:] {
			if _, exists := allDestinations[station]; exists {
				travel := dijkstra(currentPosition, station, edges)
				totalTime += travel.Cost
				currentPosition = station
				delete(allDestinations, station)
			}
		}
		trainCpy.loadedPackages = trainCpy.loadedPackages[1:]
	}

	return totalTime
}


func Permutations(items []Package) [][]Package {
    var result [][]Package
    if len(items) == 1 {
        return [][]Package{items}
    }
    for i, item := range items {
        rest := append([]Package{}, items[:i]...)
        rest = append(rest, items[i+1:]...)
        for _, perm := range Permutations(rest) {
            result = append(result, append([]Package{item}, perm...))
        }
    }
    return result
}

// Generate all possible configurations of trains and packages. Brute force. 
func generateMoves(trains []Train, packages []Package) [][]Train {
    packagePermutations := Permutations(packages)
    var allMoves [][]Train

    for _, pkgPerm := range packagePermutations {
        move := make([]Train, len(trains))
        for i, train := range trains {
            move[i] = Train{
                name:            train.name,
                currentPosition: train.currentPosition,
                capacity:        train.capacity,
                remaining:       train.capacity,
                loadedPackages:  []Package{},
            }
        }

        validConfiguration := true
        for _, pkg := range pkgPerm {
            assigned := false
            for j := range move {
                if move[j].remaining >= pkg.weight {
                    move[j].loadedPackages = append(move[j].loadedPackages, pkg)
                    move[j].remaining -= pkg.weight
                    assigned = true
                    break
                }
            }
            if !assigned {
                validConfiguration = false
                break
            }
        }

        if validConfiguration {
            allMoves = append(allMoves, move)
        }
    }
    return allMoves
}

//Given all the possible configurations, calculate the configuration with the shortest time
func findMinimalTimeConfiguration(moves [][]Train, edge []Edge) int {
	minTime := math.MaxInt64
	var bestMove []Train
	for _, move := range moves {
		maxTime := 0
		for _, train := range move {
			time := calculateTravelTime(train, edge)
			if time > maxTime {
				maxTime = time
			}
		}
		if maxTime < minTime {
			minTime = maxTime
			bestMove = move
		}
	}
	debugBestMove := [][]Train{bestMove}
	debugMoves(debugBestMove)
	return minTime
}


func main() {
	edges := []Edge{
		{"A", "B", 30},
		{"B", "C", 10},
	}
	packages := []Package{
		{"K1", 5, "A", "C"},
	}
	trains := []Train{
		{"Q1", "B", 6, 6, []Package{}},
	}

	moves := generateMoves(trains, packages)
	// debugMoves(moves)
	bestTime := findMinimalTimeConfiguration(moves, edges)

	fmt.Println("Best possible time:", bestTime)
}

func debugMoves(moves [][]Train) {
	for moveIdx, move := range moves {
		fmt.Printf("Move %d:\n", moveIdx+1)
		for _, train := range move {
			fmt.Printf("  Train: %s, Current Position: %s, Remaining Capacity: %d\n", train.name, train.currentPosition, train.remaining)
			fmt.Printf("  Loaded Packages:\n")
			for _, pkg := range train.loadedPackages {
				fmt.Printf("    Package: %s, Weight: %d, Start: %s, End: %s\n", pkg.name, pkg.weight, pkg.start, pkg.end)
			}
		}
		fmt.Println()
	}
}

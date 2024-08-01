# Given a list of trains, a list of packages, a list of stations, create a global time optimal algorithm to deliver all packages as quickly as possible.

### Approach

Brute force is the only way to find all possible configurations and assess delivery speed. Other algorithms are merely hueristics and won't always guarantee the quickest delivery of all packages. The tradeoff for precision is algorithm complexity, which shall be detailed later.

1. Initialize the graph with nodes and edges.
2. Generate all possible moves for each train at each step. `generateMoves()`
3. Using the results from (2), calculate the time taken for each configuration and store the shortest time in `calculateTravelTime()`.
4. Return the shortest time from (3).

### Complexity

`permutations()` generates all permutations of the input list of packages. The number of permutations of n packages is `n!`, thus this function is a `O(n!)` algorithm.

`generateMoves()` iterates over the trains in `O(m)` and applies itself to all `O(n!)` possible package delivery orderings. Total complexity is thus `O(m*n!)`. The only intelligence it is built in with is it checks for `validConfigurations`. A valid configuration is a configuration where all the packages can be loaded onto the train without the train getting overweight; suppose we have a configuration where the package weight is 10 and the train's capacity is 5, then this configuration is quickly rejected.

`calculateTravelTime()` calculates the total travel time for a given train configuration using Dijkstra's algorithm to find the shortest paths. Since the complexity of Dijkstra's is `O(E+VlogV)` where `E` is an edge and `V` is a vertex, the final complexity is `O(S⋅(E+VlogV))` where `S` is the number of stops we make along a route. To optimized, we have implemented a memoization for Dijkstra, such that if we've previously calculated the result for a set of inputs, we will just return it from a storage.

`main()`. Generally, the complexity is dominated by the `generateMoves()` function.

### Test case

```
stations := []string{"A", "B", "C"}
trains := []Train{
	{name: "Train1", currentPosition: "A", capacity: 100, remaining: 100, loadedPackages: []Package{}},
	{name: "Train2", currentPosition: "B", capacity: 100, remaining: 100, loadedPackages: []Package{}},
}
packages := []Package{
	{name: "Package1", weight: 10, start: "A", end: "C"},
	{name: "Package2", weight: 20, start: "B", end: "C"},
	{name: "Package3", weight: 15, start: "A", end: "B"},
}
edges := []Edge{
	{"A", "B", 30},
	{"B", "C", 10},
}
Answer: 40
```

```
stations := []string{"A", "B", "C"}
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
Answer: 70
```

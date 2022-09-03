## Game of Life
Create a function ```NextGeneration(currentPopulation []string) []string``` that takes a starting configuration for *Conway's game of life* and calculates the next generation.

## Description
[Conway’s game of life](http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) models the evolution of a population living on a two-dimensional grid.
The grid is finite and can be of any size
Every cell in the grid is either dead or alive. In the next generation, cells change their status according to the following rules.


1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbours becomes a live cell.

### Defintion of Neighborhood
In this version of the kata, the neighborhood doesn't wrap around the edges. Any cell `x` can have at most 8 neighbours `N`, as described in the following 5x5 grid example:

```
Cell with 8 neighbours:
.....
.NNN.
.NxN.
.NNN.
.....

Cell with 5 neighbours:
.NxN.
.NNN.
.....
.....
.....

Cell with 3 neighbours:
...Nx
...NN
.....
.....
.....
```

## Examples:
```
.: dead
x: alive

5x3 grid example of a single dying cell:
Gen:0    Gen:1
.....    .....
.x... => .....
.....    .....

5x3 grid example of a stabilizing pattern:
Gen:0    Gen:1    Gen:2
..x..    .xx..    .xx..
.xx.. => .xx.. => .xx..
.....    .....    .....

```


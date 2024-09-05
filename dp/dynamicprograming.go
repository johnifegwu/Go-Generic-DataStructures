package dp

import "fmt"

//const showExpectedResult = false
//const showHints = false

type Maze struct {
	values        [][]int
	computedPaths map[string]int
	rows, cols    int
}

func NewMaze(vals [][]int) *Maze {
	m := &Maze{
		values:        vals,
		computedPaths: make(map[string]int),
	}
	m.rows = len(vals)
	m.cols = len(vals[0])
	return m
}

func (m *Maze) GetCost(row, col int) int {
	return m.values[row][col]
}

func (m *Maze) IsDestination(row, col int) bool {
	return row == (m.rows-1) && col == (m.cols-1)
}

func (m *Maze) IsLastRow(row int) bool {
	return row == (m.rows - 1)
}

func (m *Maze) IsLastCol(col int) bool {
	return col == (m.cols - 1)
}

func (m *Maze) getKey(row, col, cost int) string {
	return fmt.Sprintf("%d-%d-%d", row, col, cost)
}

func (m *Maze) SetPathCount(row, col, cost, value int) {
	key := m.getKey(row, col, cost)
	m.computedPaths[key] = value
}

func (m *Maze) GetPathCount(row, col, cost int) (int, bool) {
	key := m.getKey(row, col, cost)
	val, ok := m.computedPaths[key]
	return val, ok
}

func CountPaths(maze *Maze, row int, col int, cost int) int {

	// we are out money, invalid path
	if cost < 0 {
		return 0
	}

	// check where cost is precomputed
	if val, exist := maze.GetPathCount(row, col, cost); exist {
		return val
	}

	// retreive the current cost
	currentCost := maze.GetCost(row, col)

	// end case
	if maze.IsDestination(row, col) {
		if currentCost == cost {
			return 1
		}
		return 0
	}

	// if we are on the last row we can only move right
	if maze.IsLastRow(row) {
		r := CountPaths(maze, row, col+1, cost-currentCost)
		maze.SetPathCount(row, col, cost-currentCost, r)
		return r
	}

	// if we are in thast col we can only go downward
	if maze.IsLastCol(col) {
		r := CountPaths(maze, row+1, col, cost-currentCost)
		maze.SetPathCount(row, col, cost-currentCost, r)
		return r
	}

	// otherwise we have two moves to make bith right and down
	r := CountPaths(maze, row+1, col, cost-currentCost) +
		CountPaths(maze, row, col+1, cost-currentCost)
	maze.SetPathCount(row, col, cost-currentCost, r)
	return r
}

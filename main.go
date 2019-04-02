package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	x int
	y int
}

func main() {
	// Size of the universe in m x n
	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)

	// No of generations to run
	var gen int
	fmt.Scanf("%d", &gen)

	// Get the initial live cells in the universe
	var liveCells []Cell
	var noLiveCells int
	fmt.Scanf("%d", &noLiveCells)
	for i := 0; i < noLiveCells; i++ {
		var str string
		fmt.Scanf("%s", &str)

		strs := strings.Split(str, ",")
		if len(strs) < 2 {
			os.Exit(1)
		}

		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		liveCells = append(liveCells, Cell{
			x: x,
			y: y,
		})
	}

	universe := make([][]int, m)
	for i := range universe {
		universe[i] = make([]int, n)
	}

	// Create the universe
	for _, c := range liveCells {
		universe[c.x][c.y] = 1
	}

	s := time.Now()

	for i := 0; i < gen; i++ {
		universe = ComputeGen(universe, m, n)

		extinct := true
		for k := 0; k < m; k++ {
			for l := 0; l < n; l++ {
				if universe[k][l] == 1 {
					extinct = false
					break
				}
			}
		}

		if extinct {
			gen = i + 1
			break
		}
	}

	fmt.Println("Universe for Generation ", gen)
	for k := 0; k < m; k++ {
		for l := 0; l < n; l++ {
			fmt.Printf("%d ", universe[k][l])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

	fmt.Println("time taken: ", time.Since(s))
}

func ComputeGen(univ [][]int, m, n int) [][]int {
	nextGenUniverse := make([][]int, m)
	for i := range nextGenUniverse {
		nextGenUniverse[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := getliveAdjacentCellCount(univ, i, j, m, n)
			if univ[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					nextGenUniverse[i][j] = 0
				}

				if cnt == 2 || cnt == 3 {
					nextGenUniverse[i][j] = 1
				}
			} else {
				if cnt == 3 {
					nextGenUniverse[i][j] = 1
				}
			}

		}
	}

	return nextGenUniverse
}

func getliveAdjacentCellCount(univ [][]int, i, j, m, n int) int {
	count := 0
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			x := k
			y := l
			if x < 0 {
				x = m - 1
			}

			if x > m-1 {
				x = 0
			}

			if y < 0 {
				y = n - 1
			}

			if y > n-1 {
				y = 0
			}

			if x == i && y == j {
				continue
			}

			if univ[x][y] == 1 {
				count++
			}
		}
	}

	return count
}

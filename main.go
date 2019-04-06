package main

import (
	"fmt"
	"time"
)

func main() {
	// Size of the universe in m x n
	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)

	// No of generations to run
	var gen int
	fmt.Scanf("%d", &gen)

	universe := make([][]int, m)
	for i := range universe {
		universe[i] = make([]int, n)
	}

	// Get the initial live cells in the universe
	var liveCellCount int
	fmt.Scanf("%d", &liveCellCount)
	for i := 0; i < liveCellCount; i++ {
		var r, c int
		fmt.Scanf("%d,%d", &r, &c)
		universe[r][c] = 1
	}

	s := time.Now()
	for i := 0; i < gen; i++ {
		universe = ComputeNextGen(universe, m, n)

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

func ComputeNextGen(univ [][]int, m, n int) [][]int {
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

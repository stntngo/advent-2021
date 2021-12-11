package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const _SIZE = 10

type Coordinate struct {
	x, y int
}

func (c Coordinate) Neighbors() []Coordinate {
	neighbors := make([]Coordinate, 0, 8)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			nx := c.x + dx
			ny := c.y + dy

			if nx < 0 || ny < 0 {
				continue
			}

			if nx >= _SIZE || ny >= _SIZE {
				continue
			}

			neighbors = append(neighbors, Coordinate{nx, ny})
		}
	}

	return neighbors
}

type Cavern [_SIZE][_SIZE]int

// func (c *Cavern) Equal(o Cavern) bool {
// 	for y := 0; y < _SIZE; y++ {
// 		for x := 0; x < _SIZE; x++ {
// 			if c[y][x] != o[y][x] {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

func (c *Cavern) Step() int {
	flashed := make(map[Coordinate]bool)
	lastRound := make([]Coordinate, 0, _SIZE*_SIZE)
	for y := 0; y < _SIZE; y++ {
		for x := 0; x < _SIZE; x++ {
			value := c[y][x] + 1
			c[y][x] = value

			if value > 9 {
				coord := Coordinate{x, y}
				lastRound = append(lastRound, coord)
				flashed[coord] = true
			}

		}
	}

	for len(lastRound) > 0 {
		nextRound := make([]Coordinate, 0, _SIZE*_SIZE)

		for _, coord := range lastRound {
			for _, neighbor := range coord.Neighbors() {
				value := c[neighbor.y][neighbor.x] + 1
				c[neighbor.y][neighbor.x] = value

				if value > 9 {
					if _, ok := flashed[neighbor]; ok {
						continue
					}

					nextRound = append(nextRound, neighbor)
					flashed[neighbor] = true
				}
			}
		}

		lastRound = nextRound
	}

	for y := 0; y < _SIZE; y++ {
		for x := 0; x < _SIZE; x++ {
			value := c[y][x]
			if value > 9 {
				c[y][x] = 0
			}
		}
	}

	return len(flashed)
}

func ParseCavern(r io.Reader) (Cavern, error) {
	scanner := bufio.NewScanner(r)

	var rows Cavern
	var rowIdx int
	for scanner.Scan() {
		line := scanner.Text()
		var row [_SIZE]int
		for i, str := range strings.Split(line, "") {
			num, err := strconv.Atoi(str)
			if err != nil {
				return rows, err
			}

			row[i] = num
		}

		rows[rowIdx] = row
		rowIdx++
	}

	if err := scanner.Err(); err != nil {
		return rows, err
	}

	return rows, nil
}

func main() {
	r := strings.NewReader(`4134384626
7114585257
1582536488
4865715538
5733423513
8532144181
1288614583
2248711141
6415871681
7881531438`)
	// r := strings.NewReader(`5483143223
	// 2745854711
	// 5264556173
	// 6141336146
	// 6357385478
	// 4167524645
	// 2176841721
	// 6882881134
	// 4846848554
	// 5283751526`)

	rows, err := ParseCavern(r)
	if err != nil {
		panic(err)
	}

	var totalFlashed int
	var i int
	var allflashed bool
	for !allflashed {
		flashed := rows.Step()
		if flashed == (_SIZE * _SIZE) {
			allflashed = true
			fmt.Println("All Flashed on Step", i+1)
		}

		totalFlashed += flashed
		i++
	}

	fmt.Println(totalFlashed)

}

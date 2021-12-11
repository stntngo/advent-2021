package day06

import (
	"strconv"
	"strings"
)

type LanternFish [9]uint64

func (l LanternFish) Pop() uint64 {
	var sum uint64

	for _, fish := range l {
		sum += fish
	}

	return sum
}

func SimulatePopulation(population LanternFish, days int) LanternFish {
	for day := 0; day < days; day++ {
		var next LanternFish

		next[6] = population[0]
		next[8] = population[0]

		for i := 1; i < 9; i++ {
			next[i-1] += population[i]
		}

		population = next
	}

	return population
}

func Parse(s string) (LanternFish, error) {
	var lanternfish LanternFish
	for _, num := range strings.Split(s, ",") {
		fish, err := strconv.Atoi(num)
		if err != nil {
			return lanternfish, err
		}

		lanternfish[fish]++
	}

	return lanternfish, nil
}

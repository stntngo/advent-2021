package day06

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Solution struct {
	fish LanternFish
}

func (s *Solution) Name() string {
	return "Lanternfish"
}

func (s *Solution) Load(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	fish, err := Parse(string(b))
	if err != nil {
		return err
	}

	s.fish = fish

	return nil
}

func (s *Solution) PartOne() (string, error) {
	return fmt.Sprintf("%v", SimulatePopulation(s.fish, 80).Pop()), nil
}

func (s *Solution) PartTwo() (string, error) {
	return fmt.Sprintf("%v", SimulatePopulation(s.fish, 256).Pop()), nil
}

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

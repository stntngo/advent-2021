package day12

import (
	"io"
	"strconv"
)

type Solution struct {
	caves CaveSystem
}

func (s *Solution) Name() string {
	return "Passage Pathing"
}

func (s *Solution) Load(r io.Reader) error {
	caves, err := ParseCaves(r)
	if err != nil {
		return err
	}
	s.caves = caves
	return nil
}

func (s *Solution) PartOne() (string, error) {
	return strconv.Itoa(s.caves.Start().Paths(1)), nil
}

func (s *Solution) PartTwo() (string, error) {
	return strconv.Itoa(s.caves.Start().Paths(2)), nil
}

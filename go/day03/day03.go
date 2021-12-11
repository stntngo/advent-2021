package day03

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Solution struct {
	lines [][]string
}

func (s *Solution) Name() string {
	return "Binary Diagnostic"
}

func (s *Solution) Load(r io.Reader) error {
	lines, err := Parse(r)
	if err != nil {
		return err
	}

	s.lines = lines
	return nil
}

func (s *Solution) PartOne() (string, error) {
	power, err := PowerConsumption(s.lines)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", power), nil
}
func (s *Solution) PartTwo() (string, error) {
	life, err := LifeSupport(s.lines)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", life), nil
}

func BitArrayToInt(ba []string) (uint64, error) {
	return strconv.ParseUint(strings.Join(ba, ""), 2, 64)
}

func Parse(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)

	var lines [][]string
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func Transpose(lines [][]string) [][]string {
	out := make([][]string, 0, len(lines[0]))

	for i := 0; i < len(lines[0]); i++ {
		line := make([]string, 0, len(lines))
		for _, row := range lines {
			line = append(line, row[i])
		}

		out = append(out, line)
	}

	return out
}

func Gamma(lines [][]string) []string {
	var gamma []string
	var gn, en int
	for _, line := range lines {
		counter := make(map[string]int)
		for _, value := range line {
			counter[value]++
		}

		if counter["1"] > counter["0"] {
			gn = (gn << 1) | 1
			en <<= 1
			gamma = append(gamma, "1")
		} else {
			gn <<= 1
			en = (en << 1) | 1
			gamma = append(gamma, "0")
		}

	}

	fmt.Println(gn, en)
	return gamma
}

func Invert(number []string) []string {
	var out []string

	for _, n := range number {
		if n == "1" {
			out = append(out, "0")
		} else {
			out = append(out, "1")
		}
	}

	return out
}

func Count(bits []string) (int, int) {
	var zeroes, ones int
	for _, bit := range bits {
		switch bit {
		case "0":
			zeroes++
		case "1":
			ones++
		default:
			panic("unrecognized bit")
		}
	}

	return zeroes, ones
}

func PowerConsumption(lines [][]string) (uint64, error) {
	lines = Transpose(lines)

	var gamma, epsilon uint64
	for _, line := range lines {
		zeroes, ones := Count(line)

		gamma <<= 1
		epsilon <<= 1
		if ones > zeroes {
			gamma |= 1
		} else {
			epsilon |= 1
		}

	}

	return gamma * epsilon, nil
}

func BitFilter(numbers [][]string, idx int, tgt func(int, int) bool) ([]string, error) {
	if len(numbers) == 0 {
		return nil, errors.New("no numbers left to filter")
	}

	if len(numbers) == 1 {
		return numbers[0], nil
	}

	target := "0"
	if tgt(Count(Transpose(numbers)[idx])) {
		target = "1"
	}

	var candidates [][]string
	for _, number := range numbers {
		if number[idx] == target {
			candidates = append(candidates, number)
		}
	}

	return BitFilter(candidates, idx+1, tgt)
}

func LifeSupport(lines [][]string) (uint64, error) {
	oxygenRaw, err := BitFilter(lines, 0, func(z, o int) bool { return o >= z })
	if err != nil {
		return 0, err
	}

	carbonRaw, err := BitFilter(lines, 0, func(z, o int) bool { return o < z })
	if err != nil {
		return 0, err
	}

	oxygen, err := BitArrayToInt(oxygenRaw)
	if err != nil {
		return 0, err
	}

	carbon, err := BitArrayToInt(carbonRaw)
	if err != nil {
		return 0, err
	}

	return oxygen * carbon, nil
}

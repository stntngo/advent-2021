package day01

import (
	"bufio"
	"io"
	"strconv"
)

type SonarReading []int

func (r SonarReading) DepthIncrease() int {
	var count int
	for i := 1; i < len(r); i++ {
		if r[i-1] < r[i] {
			count++
		}
	}

	return count
}

func (r SonarReading) SlidingWindow() SonarReading {
	var out []int
	for i := 2; i < len(r); i++ {
		out = append(out, r[i-2]+r[i-1]+r[i])
	}

	return out
}

func ParseSonarReading(r io.Reader) (SonarReading, error) {
	scanner := bufio.NewScanner(r)

	var reading []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		reading = append(reading, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reading, nil
}

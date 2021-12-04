package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func ParseInput(r io.Reader) (SonarReading, error) {
	scanner := bufio.NewScanner(r)

	var reading []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		reading = append(reading, i)
	}

	return reading, nil
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reading, err := ParseInput(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", reading.DepthIncrease())
	fmt.Println("Part Two:", reading.SlidingWindow().DepthIncrease())

}

// In Go 1.18 they'll be adding support for Generics. The "Sliding Window" function provides
// a great example of a situation in which generics would work well. Below is an example
// of a generic window function written in Go 1.18.
//
// func Windows[T any](v []T, size int) ([][]T, error) {
// 	if len(v) < size {
// 		return nil, errors.New("window size larger than input array")
// 	}

// 	windows := make([][]T, 0, len(v)-size-1)
// 	for size < len(v) {
// 		windows = append(windows, v[:size])
// 		v = v[1:]
// 	}

// 	return windows, nil
// }

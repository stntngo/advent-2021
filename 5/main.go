package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type LineType int

const (
	Horizontal LineType = iota + 1
	Vertical
	Diagonal
)

type Point struct {
	X, Y int
}

func ParsePoint(s string) (Point, error) {
	var point Point

	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return point, errors.New("point must be defiend as (X, Y) pair")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return point, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return point, err
	}

	point.X = x
	point.Y = y

	return point, nil
}

type Line struct {
	Start, End Point
}

func (l *Line) LineType() LineType {
	if l.Start.X == l.End.X {
		return Vertical
	}

	if l.Start.Y == l.End.Y {
		return Horizontal
	}

	return Diagonal
}

func (l *Line) Points() []Point {
	switch l.LineType() {
	case Horizontal:
		return l.horizontalLine()
	case Vertical:
		return l.verticalLine()
	case Diagonal:
		return l.diagonaLine()
	default:
		return nil
	}
}

func (l *Line) verticalLine() []Point {
	start, end := l.Start, l.End
	if start.Y > end.Y {
		start, end = end, start
	}

	points := make([]Point, 0)
	for i := start.Y; i <= end.Y; i++ {
		points = append(points, Point{start.X, i})
	}

	return points
}

func (l *Line) horizontalLine() []Point {
	start, end := l.Start, l.End
	if start.X > end.X {
		start, end = end, start
	}

	points := make([]Point, 0)
	for i := start.X; i <= end.X; i++ {
		points = append(points, Point{i, start.Y})
	}

	return points
}

func (l *Line) diagonaLine() []Point {
	start, end := l.Start, l.End
	if start.X > end.X {
		start, end = end, start
	}

	points := make([]Point, 0)

	y := start.Y
	yFunc := func(y int) int {
		return y + 1
	}

	if start.Y > end.Y {
		yFunc = func(y int) int {
			return y - 1
		}
	}
	for i := start.X; i <= end.X; i++ {
		points = append(points, Point{i, y})

		y = yFunc(y)
	}

	return points
}

func ParseLine(s string) (Line, error) {
	var line Line

	parts := strings.Split(s, " -> ")
	if len(parts) != 2 {
		return line, errors.New("lines must be defined by two points")
	}

	start, err := ParsePoint(parts[0])
	if err != nil {
		return line, err
	}

	end, err := ParsePoint(parts[1])
	if err != nil {
		return line, err
	}

	line.Start = start
	line.End = end

	return line, nil
}

func Parse(r io.Reader) ([]Line, error) {
	scanner := bufio.NewScanner(r)

	var lines []Line
	for scanner.Scan() {
		line, err := ParseLine(scanner.Text())
		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// NOTE (niels): I wrote this imperative style program first where I take in the diagonal option because
// I had shoved the ability to say whether or not I was interetested in the diagonal points of a line
// into the Points function itself. After writing the functional clojure version the "better" solution
// is letting the caller of CountHotSpots decide whether diagonal lines should be counted not by
// providing a flag parameter, but by filtering them out of the call to begin with.
func CountHotSpots(lines []Line) int {
	vents := make(map[Point]int)

	for _, line := range lines {
		for _, point := range line.Points() {
			vents[point]++
		}
	}

	var total int
	for _, count := range vents {
		if count > 1 {
			total++
		}
	}

	return total
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := Parse(f)
	if err != nil {
		panic(err)
	}

	noDiagonals := make([]Line, 0, len(lines))

	for _, line := range lines {
		if line.LineType() != Diagonal {
			noDiagonals = append(noDiagonals, line)
		}
	}

	fmt.Println("Part One:", CountHotSpots(noDiagonals))
	fmt.Println("Part Two:", CountHotSpots(lines))

}

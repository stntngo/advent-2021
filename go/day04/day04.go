package day04

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ParseRand(line string) (*RandomNumbers, error) {
	strs := strings.Split(line, ",")
	nums := make([]int, len(strs))

	for i, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		nums[i] = num
	}

	return &RandomNumbers{
		next:  nums,
		last:  -1,
		drawn: make(map[int]bool),
	}, nil
}

type RandomNumbers struct {
	next  []int
	last  int
	drawn map[int]bool
}

func (r *RandomNumbers) Draw() error {
	if len(r.next) < 1 {
		return errors.New("empty random number stream")
	}

	r.last, r.next = r.next[0], r.next[1:]
	r.drawn[r.last] = true

	return nil
}

func (r *RandomNumbers) Drawn(n int) bool {
	_, ok := r.drawn[n]
	return ok
}

func (r *RandomNumbers) WinningRow(row [5]int) bool {
	for _, number := range row {
		if !r.Drawn(number) {
			return false
		}
	}

	return true
}

func (r *RandomNumbers) WinningBoard(b Board) bool {
	for _, row := range b {
		if r.WinningRow(row) {
			return true
		}
	}

	for _, column := range b.Transpose() {
		if r.WinningRow(column) {
			return true
		}
	}

	return false
}

func (r *RandomNumbers) Score(b Board) int {
	var sum int
	for _, row := range b {
		for _, number := range row {
			if !r.Drawn(number) {
				sum += number
			}
		}
	}

	return sum * r.last
}

type Board [5][5]int

func (b Board) Transpose() Board {
	var trans [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			trans[j][i] = b[i][j]
		}
	}

	return trans
}

func WinFirst(r *RandomNumbers, boards []Board) (Board, error) {
	var empty [5][5]int
	for {
		if err := r.Draw(); err != nil {
			return empty, err
		}

		for _, board := range boards {
			if r.WinningBoard(board) {
				return board, nil
			}
		}
	}
}

func WinLast(r *RandomNumbers, boards []Board) (Board, error) {
	var empty [5][5]int
	for len(boards) > 1 {
		if err := r.Draw(); err != nil {
			return empty, err
		}

		candidates := make([]Board, 0, len(boards))
		for _, board := range boards {
			if !r.WinningBoard(board) {
				candidates = append(candidates, board)
			}
		}

		boards = candidates
	}

	last := boards[0]
	for !r.WinningBoard(last) {
		if err := r.Draw(); err != nil {
			return empty, err

		}

	}

	return last, nil
}

func ParseBoard(lines []string) (Board, error) {
	var board [5][5]int
	if len(lines) != 5 {
		return board, errors.New("board must be 5 rows")
	}

	for i := 0; i < 5; i++ {
		line := strings.FieldsFunc(lines[i], func(r rune) bool {
			return r == ' '
		})

		if len(line) != 5 {
			return board, errors.New("row must be 5 columns")
		}

		for j, str := range line {
			num, err := strconv.Atoi(str)
			if err != nil {
				return board, err
			}

			board[i][j] = num
		}
	}

	return board, nil
}

func Parse(r io.Reader) (*RandomNumbers, []Board, error) {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	rawNumbers, rawBoardLines := lines[0], lines[1:]
	rand, err := ParseRand(rawNumbers)
	if err != nil {
		return nil, nil, err
	}

	var boards []Board
	for len(rawBoardLines) > 0 {
		var rawBoard []string
		rawBoard, rawBoardLines = rawBoardLines[:5], rawBoardLines[5:]

		board, err := ParseBoard(rawBoard)
		if err != nil {
			return nil, nil, err
		}

		boards = append(boards, board)
	}

	return rand, boards, nil
}

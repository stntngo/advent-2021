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

func (r *RandomNumbers) Score(b Board) (int, bool) {
	var winner bool
	for _, row := range b {
		var match int
		for _, number := range row {
			if r.Drawn(number) {
				match++
			}
		}

		if match == len(row) {
			winner = true
			break
		}

	}

	if !winner {
		for i := 0; i < 5; i++ {
			var match int

			for _, row := range b {
				if r.Drawn(row[i]) {
					match++
				}
			}

			if match == 5 {
				winner = true
				break
			}
		}
	}

	if !winner {
		return 0, false
	}

	var sum int
	for _, row := range b {
		for _, number := range row {
			if !r.Drawn(number) {
				sum += number
			}
		}
	}

	return sum * r.last, true
}

type Board [5][5]int

func PlayGame(r *RandomNumbers, boards []Board) (Board, error) {
	for {
		if err := r.Draw(); err != nil {
			var winner [5][5]int
			return winner, err
		}

		for _, board := range boards {
			if _, won := r.Score(board); won {
				return board, nil
			}
		}
	}
}

func ParseBoard(lines [][]string) (Board, error) {
	var board [5][5]int
	if len(lines) != 5 {
		return board, errors.New("top level: board must be 5x5")
	}

	for i := 0; i < 5; i++ {
		if len(lines[i]) != 5 {
			return board, errors.New("inner: board msut be 5x5")
		}

		for j, str := range lines[i] {
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

	var rand *RandomNumbers
	var boards []Board
	buffer := make([][]string, 0, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if rand == nil {
			var err error
			rand, err = ParseRand(line)
			if err != nil {
				return nil, nil, err
			}

			continue
		}

		buffer = append(buffer, strings.FieldsFunc(line, func(r rune) bool {
			return r == ' '
		}))

		if len(buffer) == 5 {
			board, err := ParseBoard(buffer)
			if err != nil {
				return nil, nil, err
			}

			boards = append(boards, board)
			buffer = make([][]string, 0, 5)
		}
	}

	return rand, boards, nil
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	rand, boards, err := Parse(f)
	if err != nil {
		panic(err)
	}

	winner, err := PlayGame(rand, boards)
	if err != nil {
		panic(err)
	}

	score, won := rand.Score(winner)
	if !won {
		panic("returned non-winner")
	}

	fmt.Println("Part One:", score)
}

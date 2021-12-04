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

func WinLast(r *RandomNumbers, boards []Board) (Board, error) {
	for len(boards) > 1 {
		if err := r.Draw(); err != nil {
			var winner [5][5]int
			return winner, err
		}

		candidates := make([]Board, 0, len(boards))
		for _, board := range boards {
			if _, won := r.Score(board); !won {
				candidates = append(candidates, board)
			}
		}

		boards = candidates
	}

	last := boards[0]

	for {
		if err := r.Draw(); err != nil {
			var winner [5][5]int
			return winner, err

		}

		if _, won := r.Score(last); won {
			return last, nil
		}
	}
}

func ParseBoard(lines []string) (Board, error) {
	var board [5][5]int
	if len(lines) != 5 {
		return board, errors.New("top level: board must be 5x5")
	}

	for i := 0; i < 5; i++ {
		line := strings.FieldsFunc(lines[i], func(r rune) bool {
			return r == ' '
		})

		if len(line) != 5 {
			return board, errors.New("inner: board msut be 5x5")
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

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	rand, boards, err := Parse(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	winner, err := PlayGame(rand, boards)
	if err != nil {
		panic(err)
	}

	score, won := rand.Score(winner)
	if !won {
		panic("returned non-winner")
	}

	fmt.Println("Part One:", score)

	f, err = os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rand, boards, err = Parse(f)
	if err != nil {
		panic(err)
	}

	loser, err := WinLast(rand, boards)
	if err != nil {
		panic(err)
	}

	score, won = rand.Score(loser)
	if !won {
		panic("returned non-winner")
	}

	fmt.Println("Part Two:", score)
}

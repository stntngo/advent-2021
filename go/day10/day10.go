package day10

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type ErrCorruptedLine struct {
	expected, got string
}

func (e *ErrCorruptedLine) Error() string {
	return fmt.Sprintf("corrupt line: expected %q got %q", e.expected, e.got)
}

type ErrUnterminatedLine struct {
	remaining []string
}

func (e *ErrUnterminatedLine) Error() string {
	return fmt.Sprintf("unterminated line: %s remaining", strings.Join(e.remaining, ""))
}

func ParseLine(s string) error {
	var stack []string
	for _, sym := range strings.Split(s, "") {
		switch sym {
		case "(":
			stack = append(stack, ")")
		case "[":
			stack = append(stack, "]")
		case "{":
			stack = append(stack, "}")
		case "<":
			stack = append(stack, ">")
		default:
			n := len(stack) - 1

			var term string
			term, stack = stack[n], stack[:n]

			if term != sym {
				return &ErrCorruptedLine{
					expected: term,
					got:      sym,
				}
			}
		}
	}

	if len(stack) > 1 {
		return &ErrUnterminatedLine{
			remaining: stack,
		}
	}

	return nil
}

func Score(f []string) (int, int) {
	var corruptionScore int
	var autocompleteScores []int

	for _, line := range f {
		if err := ParseLine(line); err != nil {
			var corruption *ErrCorruptedLine
			if errors.As(err, &corruption) {
				switch corruption.got {
				case ")":
					corruptionScore += 3
				case "]":
					corruptionScore += 57
				case "}":
					corruptionScore += 1197
				case ">":
					corruptionScore += 25137
				default:
					panic("unknown terminator")
				}
			}

			var incomplete *ErrUnterminatedLine
			if errors.As(err, &incomplete) {
				var score int
				for i := len(incomplete.remaining) - 1; i >= 0; i-- {
					score *= 5
					switch incomplete.remaining[i] {
					case ")":
						score += 1
					case "]":
						score += 2
					case "}":
						score += 3
					case ">":
						score += 4
					default:
						panic("unknown terminator")
					}
				}

				autocompleteScores = append(autocompleteScores, score)
			}
		}
	}

	sort.Ints(autocompleteScores)
	return corruptionScore, autocompleteScores[len(autocompleteScores)/2]
}

func Lines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var out []string
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

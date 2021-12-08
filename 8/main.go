package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

var (
	_COUNTS = []int{8, 6, 8, 7, 4, 9, 7}
	_DIGITS = [][]int{
		{0, 1, 2, 4, 5, 6},    // 0
		{2, 5},                // 1
		{0, 2, 3, 4, 6},       // 2
		{0, 2, 3, 5, 6},       // 3
		{1, 2, 3, 5},          // 4
		{0, 1, 3, 5, 6},       // 5
		{0, 1, 3, 4, 5, 6},    // 6
		{0, 2, 5},             // 7
		{0, 1, 2, 3, 4, 5, 6}, // 8
		{0, 1, 2, 3, 5, 6},    // 9
	}
	_COMBINATIONS = make([][]rune, 0, 5040)
)

func init() {
	Permutations([]rune("abcdefg"), func(a []rune) {
		b := make([]rune, len(a))
		copy(b, a)
		_COMBINATIONS = append(_COMBINATIONS, b)
	})
}

// Permutations calls f with each permutation of a.
func Permutations(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func DigitSegments(pattern []rune) [][]rune {
	segments := make([][]rune, 0, 10)

	for _, d := range _DIGITS {
		digit := make([]rune, len(d))
		for i, j := range d {
			digit[i] = pattern[j]
		}

		segments = append(segments, digit)
	}

	return segments
}

func VerifySignal(segments []rune, signal Signal) bool {
	if !CountFilter(segments, signal) {
		return false
	}

	candidates := DigitSegments(segments)

	digits := make([]string, len(signal.digits))
	copy(digits, signal.digits)

	for _, candidate := range candidates {
		match := -1
		for i, digit := range digits {
			if SegmentMatch(candidate, digit) {
				match = i
				break
			}
		}

		if match < 0 {
			return false
		}

		digits[match] = digits[len(digits)-1]
		digits[len(digits)-1] = ""
		digits = digits[:len(digits)-1]
	}

	return true
}

type Signal struct {
	digits []string
	output []string
}

func (s Signal) Output() int {
	decoded := DecodeSignal(s)

	digits := DigitSegments(decoded)

	var final int
	for i, output := range s.output {
		d := DecodeOutput(digits, output)
		final += int(float64(d) * math.Pow(10, float64(len(s.output)-i-1)))
	}

	return final

}

func ParseSignal(s string) Signal {
	var signal Signal

	parts := strings.Split(s, " | ")
	for _, digit := range strings.Split(parts[0], " ") {
		signal.digits = append(signal.digits, digit)
	}

	for _, output := range strings.Split(parts[1], " ") {
		signal.output = append(signal.output, output)
	}

	return signal
}

func Parse(r io.Reader) ([]Signal, error) {
	scanner := bufio.NewScanner(r)

	var signals []Signal

	for scanner.Scan() {
		signal := ParseSignal(scanner.Text())

		signals = append(signals, signal)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return signals, nil
}

// Treat this like a "make change" problem when there are multiple possible values then
// recurse down with those options each specified.
func CountFilter(mapping []rune, signal Signal) bool {
	for i, c := range _COUNTS {

		var count int
		for _, digit := range signal.digits {
			if strings.ContainsRune(digit, mapping[i]) {
				count++
			}
		}

		if count != c {
			// fmt.Println(string(mapping), i, count, c)
			return false
		}
	}

	return true

}

func SegmentMatch(digit []rune, signal string) bool {
	if len(digit) != len(signal) {
		return false
	}

	for _, r := range digit {
		if !strings.ContainsRune(signal, r) {
			return false
		}
	}

	return true
}

func DecodeSignal(signal Signal) []rune {
	for _, candidate := range _COMBINATIONS {
		if VerifySignal(candidate, signal) {
			return candidate
		}
	}

	panic("unreachable")
}

func DecodeOutput(rules [][]rune, output string) int {
	for d, rule := range rules {
		if SegmentMatch(rule, output) {
			return d
		}
	}

	panic("unreachable")
}

func EasyDigitCount(signals []Signal) int {
	var count int
	for _, signal := range signals {
		for _, i := range []int{2, 3, 4, 7} {
			for _, output := range signal.output {
				if len(output) == i {
					count++
				}
			}

		}
	}

	return count
}

func SignalOutputSum(signals []Signal) int {
	var sum int
	for _, signal := range signals {
		sum += signal.Output()
	}

	return sum
}

func main() {
	r, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	signals, err := Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", EasyDigitCount(signals))
	fmt.Println("Part Two:", SignalOutputSum(signals))
}

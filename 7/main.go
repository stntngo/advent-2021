package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type CostFunction func(int) float64

func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// I've already implemented the boring closed form solution over in the
// Clojure side of things. So over here I'm having a bit more fun with
// things and choosing to instead implement a silly line search in which
// the step size is arbitrarily halved each time, leading ultimately
// to a final 1-lookahead greeddy algorithm searching for the minimum.
func SillyLineSearch(cost CostFunction, guess, step int) int {
	reference := cost(guess)
	lower := cost(guess - step)
	higher := cost(guess + step)

	switch {
	case step == 1 && (reference < lower && reference < higher):
		return int(reference)
	case lower < reference:
		return SillyLineSearch(cost, guess-step, Max(1, step/2))
	case higher < reference:
		return SillyLineSearch(cost, guess+step, Max(1, step/2))
	default:
		return SillyLineSearch(cost, guess, Max(1, step/2))
	}
}

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	var ints []int

	for _, str := range strings.Split(string(b), ",") {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			panic(err)
		}

		ints = append(ints, num)
	}

	l1 := func(x int) float64 {
		var cost float64

		for _, num := range ints {
			cost += math.Abs(float64(num - x))
		}

		return cost
	}

	l2 := func(x int) float64 {
		var cost float64

		for _, num := range ints {
			val := math.Abs(float64(num - x))
			cost += val * (val + 1) / 2
		}

		return cost
	}

	sort.Ints(ints)
	mid := (ints[0] + ints[len(ints)-1]) / 2
	fmt.Println("Part One:", int(SillyLineSearch(l1, mid, 256)))
	fmt.Println("Part Two:", int(SillyLineSearch(l2, mid, 256)))

}

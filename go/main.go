package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/stntngo/advent-2021/go/day01"
)

type Solution interface {
	Name() string
	Load(io.Reader) error
	PartOne() (string, error)
	PartTwo() (string, error)
}

//go:embed input
var inputs embed.FS

var solutions = []Solution{
	&day01.Solution{},
}

func main() {
	table := make([][]string, 0, len(solutions))
	for i, sol := range solutions {
		tstart := time.Now()
		f, err := inputs.Open(fmt.Sprintf("input/day-%02d", i+1))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := sol.Load(f); err != nil {
			panic(err)
		}

		pone, err := sol.PartOne()
		if err != nil {
			panic(err)
		}

		ptwo, err := sol.PartTwo()
		if err != nil {
			panic(err)
		}
		tend := time.Now()

		row := []string{
			fmt.Sprintf("Day %v", i+1),
			sol.Name(),
			pone,
			ptwo,
			fmt.Sprintf("%s", tend.Sub(tstart)),
		}

		table = append(table, row)

	}

	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"Day", "Name", "Part One", "Part Two", "Duration"})
	w.SetBorder(false)
	w.AppendBulk(table)
	w.Render()
}

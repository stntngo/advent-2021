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

type Direction uint

const (
	Forward Direction = iota + 1
	Down
	Up
)

type Command struct {
	Dir      Direction
	Distance int
}

func ParseCommand(str string) (Command, error) {
	parts := strings.Split(str, " ")
	if len(parts) != 2 {
		return Command{}, errors.New("command format '[direction] [distance]'")
	}

	var dir Direction

	switch parts[0] {
	case "forward":
		dir = Forward
	case "down":
		dir = Down
	case "up":
		dir = Up
	default:
		return Command{}, errors.New("unknown direction")
	}

	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return Command{}, fmt.Errorf("can't parse distance: %w", err)
	}

	return Command{
		Dir:      dir,
		Distance: i,
	}, nil
}

type Position struct {
	x   int
	y   int
	aim int
}

func (p *Position) Move(c Command) {
	switch c.Dir {
	case Forward:
		p.x += c.Distance
	case Down:
		p.y += c.Distance
	case Up:
		p.y -= c.Distance
	}
}

func (p *Position) Aim(c Command) {
	switch c.Dir {
	case Forward:
		p.x += c.Distance
		p.y += p.aim * c.Distance
	case Down:
		p.aim += c.Distance
	case Up:
		p.aim -= c.Distance
	}
}

func (p Position) Vector() int {
	return p.x * p.y
}

func Vector(commands []Command) int {
	var p Position

	for _, command := range commands {
		p.Move(command)
	}

	return p.Vector()
}

func AimVector(commands []Command) int {
	var p Position

	for _, command := range commands {
		p.Aim(command)
	}

	return p.Vector()
}

func ParseCommands(r io.Reader) ([]Command, error) {
	var commands []Command

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		command, err := ParseCommand(scanner.Text())
		if err != nil {
			return nil, err
		}

		commands = append(commands, command)
	}

	return commands, nil
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	commands, err := ParseCommands(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part One:", Vector(commands))
	fmt.Println("Part Two:", AimVector(commands))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type frame interface {
	NewValue(*cpu) error
	NextAddress(cpu, int) (int, error)
}

type valueFrame struct {
	instruction string
	register    string
}

func (f valueFrame) NewValue(c *cpu) error {
	current, err := c.Get(f.register)
	if err != nil {
		return errors.Wrap(err, "failed to get register value")
	}

	switch f.instruction {
	case "inc":
		current++
	case "tpl":
		current *= 3
	case "hlf":
		current /= 2
	default:
		return fmt.Errorf("Unknown instruction: %s", f.instruction)
	}

	err = c.Store(f.register, current)
	if err != nil {
		return errors.Wrap(err, "valueFrame failed to store")
	}
	return nil
}

func (f valueFrame) NextAddress(c cpu, current int) (int, error) {
	return current + 1, nil
}

type jumpFrame struct {
	instruction string
	register    string
	offset      int
}

func (f jumpFrame) NewValue(c *cpu) error {
	return nil
}

func (f jumpFrame) NextAddress(c cpu, current int) (int, error) {
	var next int
	switch f.instruction {
	case "jmp":
		next = current + f.offset
	case "jie":
		value, err := c.Get(f.register)
		if err != nil {
			return 0, errors.Wrap(err, "jie failed to get register value")
		}
		if value%2 == 0 {
			next = current + f.offset
		} else {
			next = current + 1
		}
	case "jio":
		value, err := c.Get(f.register)
		if err != nil {
			return 0, errors.Wrap(err, "jio failed to get register value")
		}
		if value == 1 {
			next = current + f.offset
		} else {
			next = current + 1
		}
	default:
		return 0, fmt.Errorf("Unknown jump instruction: %s", f.instruction)
	}
	return next, nil
}

type cpu struct {
	a int
	b int
}

func (c *cpu) Store(register string, value int) error {
	switch register {
	case "a":
		c.a = value
	case "b":
		c.b = value
	default:
		return fmt.Errorf("Unknown register: %s", register)
	}

	return nil
}

func (c cpu) Get(register string) (int, error) {
	var value int
	switch register {
	case "a":
		value = c.a
	case "b":
		value = c.b
	default:
		return 0, fmt.Errorf("Unknown register: %s", register)
	}

	return value, nil
}

func main() {

	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	var stack = []frame{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		switch parts[0] {
		case "hlf", "tpl", "inc":
			stack = append(stack, valueFrame{instruction: parts[0], register: parts[1]})
		case "jmp":
			offset, _ := strconv.Atoi(parts[1])
			stack = append(stack, jumpFrame{instruction: parts[0], offset: offset})
		case "jio", "jie":
			offset, _ := strconv.Atoi(parts[2])
			register := parts[1][0:1]
			stack = append(stack, jumpFrame{instruction: parts[0],
				offset:   offset,
				register: register,
			})

		}
	}

	c := cpu{a: 1}

	for i := 0; i < len(stack); {
		frame := stack[i]
		fmt.Printf("frame = %+v\n", frame)

		err = frame.NewValue(&c)
		if err != nil {
			panic(err)
		}

		i, err = frame.NextAddress(c, i)
		if err != nil {
			panic(err)
		}
		fmt.Printf("c = %+v\n", c)
		fmt.Printf("i = %+v\n", i)
	}

	fmt.Printf("c = %+v\n", c)
}

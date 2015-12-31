// Package main provides a program to convert a circuit diagram
// provided in circuit.txt in the current directory of the format
// as x AND y -> z means to connect wires x and y to an AND gate, and then
//  connect its output to wire z.

//  For example:

// -  123 -> x means that the signal 123 is provided to wire x.
// -  x AND y -> z means that the bitwise AND of wire x and wire y is provided
//    to wire z.
// -  p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and
//    then provided to wire q.
// -  NOT e -> f means that the bitwise complement of the value from wire e is
//    provided to wire f.

// Other possible gates include OR (bitwise OR) and RSHIFT (right-shift).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type wire interface {
	getSignal() uint16
	trace()
}

type binaryWire struct {
	operation func(uint16, uint16) uint16
	lhs       string
	rhs       string
}

type shiftWire struct {
	operation func(uint16, uint16) uint16
	lhs       string
	amount    uint16
}

type partialWire struct {
	operation func(uint16, uint16) uint16
	lhs       uint16
	rhs       string
}

type valueWire struct {
	signal uint16
}

func (w valueWire) trace() {
	fmt.Println("END at", w.signal, "\n")
}

func (w binaryWire) trace() {
	if w.rhs != "" {
		fmt.Println(w.lhs, w.rhs)
	}
	fmt.Println(w.lhs)
	wires[w.lhs].trace()
	if w.rhs != "" {
		fmt.Println(2, w.rhs)
		wires[w.rhs].trace()
	}
}

func (w shiftWire) trace() {
	fmt.Println(w.lhs)
	wires[w.lhs].trace()
}

func (w partialWire) trace() {
	fmt.Println(w.rhs)
	wires[w.rhs].trace()
}

func (w valueWire) getSignal() uint16 {
	return w.signal
}

func (w binaryWire) getSignal() uint16 {
	left := wires[w.lhs].getSignal()
	wires[w.lhs] = valueWire{signal: left}
	right := uint16(0)
	if w.rhs != "" {
		right = wires[w.rhs].getSignal()
		wires[w.rhs] = valueWire{signal: right}
	}
	return w.operation(left, right)
}

func (w shiftWire) getSignal() uint16 {
	left := wires[w.lhs].getSignal()
	wires[w.lhs] = valueWire{signal: left}
	return w.operation(left, w.amount)
}

func (w partialWire) getSignal() uint16 {
	right := wires[w.rhs].getSignal()
	wires[w.rhs] = valueWire{signal: right}
	return w.operation(w.lhs, right)
}

var wires = make(map[string]wire)

func main() {
	file, err := os.Open("circuit.txt")
	if err != nil {
		fmt.Println("Error opening circuit.txt", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		name := parts[1]
		wire := newWire(parts[0])

		wires[name] = wire
	}

	fmt.Println("Bobby, wire a has the value", wires["a"].getSignal(), "on it.")
}

func newWire(s string) wire {
	parts := strings.Split(s, " ")

	switch {
	case len(parts) == 1:
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			return binaryWire{operation: opValue, lhs: parts[0]}
		} else {
			return valueWire{signal: uint16(value)}
		}
	case len(parts) == 2:
		return binaryWire{operation: opNot, lhs: parts[1], rhs: ""}
	case parts[1] == "AND":
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			return binaryWire{operation: opAnd, lhs: parts[0], rhs: parts[2]}
		} else {
			return partialWire{operation: opAnd, lhs: uint16(value), rhs: parts[2]}
		}
		value, err = strconv.Atoi(parts[2])
		if err != nil {
			return binaryWire{operation: opAnd, lhs: parts[0], rhs: parts[2]}
		} else {
			return partialWire{operation: opAnd, lhs: uint16(value), rhs: parts[2]}
		}
	case parts[1] == "OR":
		return binaryWire{operation: opOr, lhs: parts[0], rhs: parts[2]}
	case parts[1] == "LSHIFT":
		value, err := strconv.Atoi(parts[2])
		if err == nil {
			return shiftWire{operation: opLShift, lhs: parts[0], amount: uint16(value)}
		}
	case parts[1] == "RSHIFT":
		value, err := strconv.Atoi(parts[2])
		if err == nil {
			return shiftWire{operation: opRShift, lhs: parts[0], amount: uint16(value)}
		}
	default:
		fmt.Println("Unknown operation", parts[1])
		panic(parts)
	}
	return valueWire{}
}

func opNot(lhs, _ uint16) uint16 {
	return ^lhs
}

func opAnd(lhs, rhs uint16) uint16 {
	return lhs & rhs
}

func opOr(lhs, rhs uint16) uint16 {
	return lhs | rhs
}

func opLShift(lhs, rhs uint16) uint16 {
	return lhs << rhs
}

func opRShift(lhs, rhs uint16) uint16 {
	return lhs >> rhs
}

func opValue(lhs, _ uint16) uint16 {
	return lhs
}

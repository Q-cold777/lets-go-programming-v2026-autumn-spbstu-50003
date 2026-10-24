package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemp = 15
	maxTemp = 30
)

type tokenReader struct {
	scanner *bufio.Scanner
}

func (r *tokenReader) next() (string, bool) {
	if !r.scanner.Scan() {
		return "", false
	}

	return r.scanner.Text(), true
}

type Thermostat struct {
	low  int
	high int
}

func NewThermostat() *Thermostat {
	return &Thermostat{
		low:  minTemp,
		high: maxTemp,
	}
}

func (t *Thermostat) ApplyConstraint(operation string, value int) {
	switch operation {
	case ">=":
		if value > t.low {
			t.low = value
		}
	case "<=":
		if value < t.high {
			t.high = value
		}
	}
}

func (t *Thermostat) GetOptimal() int {
	if t.low <= t.high {
		return t.low
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)

	defer func() { _ = out.Flush() }()

	reader := &tokenReader{scanner: scanner}

	token, hasNext := reader.next()
	if !hasNext {
		return
	}

	numDays, err := strconv.Atoi(token)
	if err != nil {
		return
	}

	for range numDays {
		if !processDay(reader, out) {
			return
		}
	}
}

func processDay(reader *tokenReader, out *bufio.Writer) bool {
	token, hasNext := reader.next()
	if !hasNext {
		return false
	}

	numConstraints, err := strconv.Atoi(token)
	if err != nil {
		return false
	}

	thermostat := NewThermostat()

	for range numConstraints {
		if !processConstraint(reader, out, thermostat) {
			return false
		}
	}

	return true
}

func processConstraint(reader *tokenReader, out *bufio.Writer, thermo *Thermostat) bool {
	operation, hasNext := reader.next()
	if !hasNext {
		return false
	}

	valueStr := strings.TrimLeft(operation, "<>=")
	operation = operation[:len(operation)-len(valueStr)]

	if valueStr == "" {
		valueStr, hasNext = reader.next()
		if !hasNext {
			return false
		}
	}

	val, err := strconv.Atoi(valueStr)
	if err != nil {
		return false
	}

	thermo.ApplyConstraint(operation, val)

	_, _ = fmt.Fprintln(out, thermo.GetOptimal())

	return true
}

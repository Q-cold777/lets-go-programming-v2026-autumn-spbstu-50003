package main

import (
	"fmt"
)

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
)

type ClimateControl struct {
	minTemp int
	maxTemp int
}

func NewClimateControl() ClimateControl {
	return ClimateControl{
		minTemp: DefaultMinTemp,
		maxTemp: DefaultMaxTemp,
	}
}

func (c ClimateControl) WithMin(val int) ClimateControl {
	if val > c.minTemp {
		c.minTemp = val
	}

	return c
}

func (c ClimateControl) WithMax(val int) ClimateControl {
	if val < c.maxTemp {
		c.maxTemp = val
	}

	return c
}

func (c ClimateControl) GetOptimal() int {
	if c.minTemp <= c.maxTemp {
		return c.minTemp
	}

	return -1
}

func main() {
	var departmentCount int
	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Ошибка чтения количества отделов:", err)

		return
	}

	for range departmentCount {
		var employeeCount int
		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Ошибка чтения количества сотрудников:", err)

			return
		}

		climate := NewClimateControl()

		for range employeeCount {
			var (
				operator string
				val      int
			)

			if _, err := fmt.Scan(&operator, &val); err != nil {
				fmt.Println("Ошибка чтения параметров сотрудника:", err)

				return
			}

			if operator == ">=" {
				climate = climate.WithMin(val)
			} else if operator == "<=" {
				climate = climate.WithMax(val)
			}

			fmt.Println(climate.GetOptimal())
		}
	}
}

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

	for dept := 0; dept < departmentCount; dept++ {
		var employeeCount int
		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Ошибка чтения количества сотрудников:", err)
			return
		}

		climate := NewClimateControl()

		for emp := 0; emp < employeeCount; emp++ {
			var (
				op  string
				val int
			)

			if _, err := fmt.Scan(&op, &val); err != nil {
				fmt.Println("Ошибка чтения параметров сотрудника:", err)
				return
			}

			if op == ">=" {
				climate = climate.WithMin(val)
			} else if op == "<=" {
				climate = climate.WithMax(val)
			}

			// Вывод результата
			fmt.Println(climate.GetOptimal())
		}
	}
}
package main

import (
	"errors"
	"fmt"
)

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
	MinCount       = 1
	MaxCount       = 1000
)

var ErrUnknownOperation = errors.New("unknown operation")

type TemperatureManager struct {
	minTemp int
	maxTemp int
}

func NewTemperatureManager() *TemperatureManager {
	return &TemperatureManager{
		minTemp: DefaultMinTemp,
		maxTemp: DefaultMaxTemp,
	}
}

func (manager *TemperatureManager) AddConstraint(operation string, value int) error {
	switch operation {
	case ">=":
		if value > manager.minTemp {
			manager.minTemp = value
		}
	case "<=":
		if value < manager.maxTemp {
			manager.maxTemp = value
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnknownOperation, operation)
	}

	return nil
}

func (manager *TemperatureManager) GetOptimalTemperature() int {
	if manager.minTemp <= manager.maxTemp {
		return manager.minTemp
	}

	return -1
}

func processEmployees(employeesCount int) error {
	manager := NewTemperatureManager()

	for range employeesCount {
		var (
			operation string
			val       int
		)

		if _, err := fmt.Scan(&operation, &val); err != nil {
			return fmt.Errorf("error while reading employees data: %w", err)
		}

		if err := manager.AddConstraint(operation, val); err != nil {
			return fmt.Errorf("error processing constraint: %w", err)
		}

		fmt.Println(manager.GetOptimalTemperature())
	}

	return nil
}

func main() {
	var departmentNumber int
	if _, err := fmt.Scan(&departmentNumber); err != nil {
		fmt.Printf("error while reading number of departments: %v\n", err)

		return
	}

	if departmentNumber < MinCount || departmentNumber > MaxCount {
		fmt.Println("incorrect number of departments")

		return
	}

	for range departmentNumber {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Printf("error while reading number of employees: %v\n", err)

			return
		}

		if employeesCount < MinCount || employeesCount > MaxCount {
			fmt.Println("incorrect employees number")

			return
		}

		if err := processEmployees(employeesCount); err != nil {
			fmt.Println(err)

			return
		}
	}
}

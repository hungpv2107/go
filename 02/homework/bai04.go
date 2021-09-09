package main

import (
	"fmt"
	"sort"
)

/*
Bài tập 4
*/

type employee struct {
	name              string
	salaryCoefficient float64
	allowance         float64
}

func (e employee) string() string {
	return fmt.Sprintf("Name: %s - Salary Coefficient: %f - Allowance: %f ", e.name, e.salaryCoefficient, e.allowance)
}

func (e employee) salary() string {
	return fmt.Sprintf("Name: %s - Salary: %f", e.name, e.getSalary())
}

func (e employee) getSalary() float64 {
	return e.salaryCoefficient*1500000 + e.allowance
}

func main() {
	employees := []employee{
		{
			name:              "DDD",
			salaryCoefficient: 1.0,
			allowance:         3000000.0,
		},
		{
			name:              "BBB",
			salaryCoefficient: 2.0,
			allowance:         1000000.0,
		},
		{
			name:              "CCC",
			salaryCoefficient: 3.0,
			allowance:         2000000.0,
		},
		{
			name:              "AAA",
			salaryCoefficient: 1.0,
			allowance:         3000000.0,
		},
	}

	// Sắp xếp theo tên tăng dần
	sortEmployeesByName(employees)
	fmt.Println("-- Sắp xếp theo tên tăng dần--")
	for _, ememployee := range employees {
		fmt.Println(ememployee.string())
	}

	// Sắp xếp theo salary giảm dần
	fmt.Println("\n-- Sắp xếp theo salary giảm dần--")
	sortEmployeesBySalary(employees)
	for _, ememployee := range employees {
		fmt.Println(ememployee.salary())
	}

	// Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
	fmt.Println("\n-- Danh sách những nhân viên có mức lương lớn thứ 2--")
	secondHighestSalary := getSecondHighestSalary(employees)
	for _, ememployee := range employees {
		if ememployee.getSalary() == secondHighestSalary {
			fmt.Println(ememployee.salary())
		}
	}
}

/*
Lấy ra mức lương cao thứ 2
*/
func getSecondHighestSalary(employees []employee) (secondHighestSalary float64) {
	highestSalary := 0.0
	secondHighestSalary = 0.0
	for _, employee := range employees {
		salary := employee.getSalary()
		if highestSalary < salary {
			secondHighestSalary = highestSalary
			highestSalary = salary
		} else if secondHighestSalary < salary {
			secondHighestSalary = salary
		}
	}
	return
}

/*
Sắp xếp danh sách nhân viên theo tên tăng dần
*/
func sortEmployeesByName(employees []employee) {
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].name < employees[j].name
	})
}

/*
Sắp xếp danh sách nhân viên theo lương giảm dần
*/
func sortEmployeesBySalary(employees []employee) {
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].getSalary() > employees[j].getSalary()
	})
}

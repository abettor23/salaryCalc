package main

import (
	"fmt"
)

func main() {
	fmt.Println("Программа вычисления зарплаты")

	amountOfStaff := 3
	employerOneSalary := 10000
	employerTwoSalary := 50000
	employerThreeSalary := 40000

	var highestSalary int
	var smallestSalary int

	if employerOneSalary > employerTwoSalary {
		if employerOneSalary > employerThreeSalary {
			highestSalary = employerOneSalary
			if employerTwoSalary > employerThreeSalary {
				smallestSalary = employerThreeSalary
			}
		} else {
			smallestSalary = employerTwoSalary
		}
	} else if employerTwoSalary > employerThreeSalary {
		highestSalary = employerTwoSalary
		if employerOneSalary > employerThreeSalary {
			smallestSalary = employerThreeSalary
		} else {
			smallestSalary = employerOneSalary
		}
	} else if employerThreeSalary > employerTwoSalary {
		highestSalary = employerThreeSalary
		smallestSalary = employerOneSalary
	}

	salaryDifference := highestSalary - smallestSalary
	fmt.Print("Разница между самой высокой зарплатой и самой низкой: ", salaryDifference, " руб.\n")

	averageSalary := (employerOneSalary + employerTwoSalary + employerThreeSalary) / amountOfStaff
	fmt.Print("Средний показатель зарплаты в отделе: ", averageSalary, " руб.\n")
}

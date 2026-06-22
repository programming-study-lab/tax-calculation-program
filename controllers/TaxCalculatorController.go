package controllers

import (
	"cal-tex/models"
	"fmt"
)

func TaxCalculatorController(taxCalculator models.TaxCalculator) float64 {
	var taxResult = 0.0
	var tax = 0.0
	const freeTax = 60000
	var totalIncome = float64(taxCalculator.TotalIncome) - freeTax

	fmt.Printf("\n{totalIncome= %.2f, taxResult = %0.2f} \n", totalIncome, taxResult)

	tax = 0.1
	if totalIncome >= 150001 {
		taxResult += (500000 - 150000) * tax

		totalIncome -= 500000

		if totalIncome <= 500000 {
			tax = 0.15
			taxResult += totalIncome * tax
		}
		fmt.Printf("\n{1 totalIncome= %.2f, taxResult = %0.2f} \n", totalIncome, taxResult)
	}

	tax = 0.15
	if totalIncome >= 500001 {
		taxResult += (1000000 - 500000) * tax

		totalIncome -= 1000000

		if totalIncome <= 1000000 {
			taxResult += totalIncome * tax
		}
		fmt.Printf("\n{2 totalIncome= %.2f, taxResult = %0.2f} \n", totalIncome, taxResult)
	}

	tax = 0.20
	if totalIncome >= 1000000 {
		taxResult += (2000000 - 1000000) * tax

		totalIncome -= 2000000

		if totalIncome <= 2000000 {
			taxResult += totalIncome * tax
		}
		fmt.Printf("\n{2 totalIncome= %.2f, taxResult = %0.2f} \n", totalIncome, taxResult)
	}

	tax = 0.35
	if totalIncome >= 2000000 {
		taxResult += totalIncome * tax
		fmt.Printf("\n{2 totalIncome= %.2f, taxResult = %0.2f} \n", totalIncome, taxResult)
	}

	if taxCalculator.Wht != 0 {
		taxResult = taxResult - float64(taxCalculator.Wht)
	}

	fmt.Printf("\n{3 totalIncome= %.2f, taxResult = %0.2f} \n", totalIncome, taxResult)

	return taxResult
}

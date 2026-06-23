package controllers

import (
	"cal-tex/models"
)

func TaxCalculatorController(taxCalculator models.TaxCalculator) models.TaxCalculatorResponse {

	var taxResult = 0.0
	var tax = 0.0
	const freeTax = 60000
	var totalIncome = 0.0
	var totalAmount = 0.0

	if len(taxCalculator.Allowances) != 0 {
		var taxAllowances = taxCalculator.Allowances

		for i := 0; i < len(taxCalculator.Allowances); i++ {
			totalAmount += float64(taxAllowances[i].Amount)
		}

		// if totalAmount <= 100000 {
		// 	totalIncome = float64(taxCalculator.TotalIncome) - 100000

		// 	// if totalAmount > 0 {
		// 	// 	totalIncome += totalAmount
		// 	// }
		// }

	}

	if totalAmount > 100000.0 {
		totalIncome = float64(taxCalculator.TotalIncome) - freeTax - totalAmount
		// totalIncome = float64(taxCalculator.TotalIncome) - freeTax + (totalAmount - 100000.0 - 100000.0)
		// fmt.Printf("\naaaaa: %0.2f, %0.2f\n", totalIncome, totalAmount)
	} else if totalAmount <= 100000.0 {
		totalIncome = float64(taxCalculator.TotalIncome) - freeTax - totalAmount
	}

	taxLevel := []models.TaxLevel{}
	taxLevel = append(taxLevel,
		models.TaxLevel{
			Level: "0-150,000",
			Tax:   int(taxResult),
		},
		models.TaxLevel{
			Level: "150,001-500,000",
			Tax:   int(taxResult),
		},
		models.TaxLevel{
			Level: "500,001-1,000,000",
			Tax:   int(taxResult),
		},
		models.TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax:   int(taxResult),
		},
		models.TaxLevel{
			Level: "2,000,001 ขึ้นไป",
			Tax:   int(taxResult),
		},
	)

	if totalIncome >= 0 {
		tax = 0.0
		taxResult = 0

		taxLevel[0] = models.TaxLevel{
			Level: "0-150,000",
			Tax:   int(taxResult),
		}
	}
	if totalIncome >= 150001 {
		tax = 0.1
		calTax := (500000 - 150000) * tax
		taxResult += calTax
		taxLevel[1] = models.TaxLevel{
			Level: "150,001-500,000",
			Tax:   int(calTax),
		}
	}
	if totalIncome >= 500001 {
		tax = 0.15
		calTax := (1000000 - 500000) * tax
		taxResult += calTax

		taxLevel[2] = models.TaxLevel{
			Level: "500,001-1,000,000",
			Tax:   int(calTax),
		}
	}
	if totalIncome >= 1000001 {
		tax = 0.2
		calTax := (2000000 - 1000000) * tax
		taxResult += calTax

		taxLevel[3] = models.TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax:   int(calTax),
		}
	}
	if totalIncome >= 2000001 {
		tax = 0.35
		calTax := totalIncome * tax
		taxResult += calTax

		taxLevel[4] = models.TaxLevel{
			Level: "2,000,001 ขึ้นไป",
			Tax:   int(calTax),
		}
	}

	if tax == 0.1 {
		taxResult -= (500000 - 150000) * tax
		calTax := (totalIncome - 150000) * tax
		taxResult += calTax

		taxLevel[1] = models.TaxLevel{
			Level: "150,001-500,000",
			Tax:   int(calTax),
		}
	} else if tax == 0.15 {
		taxResult -= (1000000 - 500000) * tax
		calTax := (totalIncome - 500000) * tax
		taxResult += calTax

		taxLevel[2] = models.TaxLevel{
			Level: "500,001-1,000,000",
			Tax:   int(calTax),
		}
	} else if tax == 0.2 {
		taxResult -= (2000000 - 1000000) * tax
		calTax := (totalIncome - 1000000) * tax
		taxResult += calTax

		taxLevel[3] = models.TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax:   int(calTax),
		}
	}

	if taxCalculator.Wht != 0 {
		taxResult = taxResult - float64(taxCalculator.Wht)
	}

	return models.TaxCalculatorResponse{
		Tax:      int(taxResult),
		TaxLevel: taxLevel,
	}
}

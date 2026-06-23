package test

import (
	"cal-tex/models"
	"cal-tex/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnTest(ctx *gin.Context) {

	allowances := []models.Allowances{}

	allowances = append(allowances, models.Allowances{
		AllowancesType: "donation",
		Amount:         0,
	})

	taxCal := models.TaxCalculator{
		TotalIncome: 1200000,
		Wht:         0,
		Allowances:  allowances,
	}

	ctx.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"status":  true,
			"message": "success",
			"data":    services.TaxCalculatorService(taxCal),
		},
	)
}

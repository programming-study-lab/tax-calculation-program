package controllers

import (
	"cal-tex/models"
	"cal-tex/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaxController(ctx *gin.Context) {

	taxCalculator := models.TaxCalculator{}

	if err := ctx.BindJSON(&taxCalculator); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error,
			},
		)
	} else {
		if taxCalculator.TotalIncome < 0 || taxCalculator.Wht < 0 {

			statusCode := http.StatusBadRequest
			statusText := http.StatusText(statusCode)

			ctx.AbortWithStatusJSON(
				statusCode,
				gin.H{
					"status":  false,
					"message": "totalIncome must be a positive number",
					"data": gin.H{
						"statusCode": statusCode,
						"statusText": statusText,
					},
				},
			)
		} else if taxCalculator.TotalIncome < taxCalculator.Wht {

			statusCode := http.StatusBadRequest
			statusText := http.StatusText(statusCode)

			ctx.AbortWithStatusJSON(
				statusCode,
				gin.H{
					"status":  false,
					"message": "wht cannot be greater than totalIncome",
					"data": gin.H{
						"statusCode": statusCode,
						"statusText": statusText,
					},
				},
			)

		} else {
			taxCalResponse := services.TaxCalculatorService(taxCalculator)

			ctx.AbortWithStatusJSON(
				http.StatusOK,
				&taxCalResponse,
			)
		}

	}

}

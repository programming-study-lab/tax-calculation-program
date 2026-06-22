package main

import (
	"cal-tex/controllers"
	"cal-tex/models"
	"cal-tex/test"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// api := r.Group("/tax")

	r.GET("test", test.OnTest)
	r.POST("/tax/calculations", func(ctx *gin.Context) {

		taxCalculator := models.TaxCalculator{}

		if err := ctx.Bind(&taxCalculator); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"status":  false,
					"message": "error",
					"data":    err.Error,
				},
			)
		} else {
			taxCal := controllers.TaxCalculatorController(taxCalculator)

			// ctx.AbortWithStatusJSON(
			// 	http.StatusOK,
			// 	gin.H{
			// 		"tax": &taxCal,
			// 	},
			// )
			ctx.AbortWithStatusJSON(
				http.StatusOK,
				&taxCal,
			)

		}

	})

	r.Run(":5000")
}

package main

import (
	"cal-tex/test"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// api := r.Group("/tax")

	r.GET("test", func(ctx *gin.Context) {

		test.OnTest(ctx)

		// ctx.AbortWithStatusJSON(
		// 	http.StatusOK,
		// 	gin.H{
		// 		"status":  true,
		// 		"message": "success",
		// 		"data":    "test",
		// 	},
		// )
	})

	r.Run(":5000")
}

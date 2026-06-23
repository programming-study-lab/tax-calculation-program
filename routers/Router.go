package routers

import (
	"cal-tex/controllers"
	"cal-tex/test"

	"github.com/gin-gonic/gin"
)

func Tax(r *gin.Engine) {
	api := r.Group("tax")
	{
		api.GET("/test", test.OnTest)
		api.POST("/calculations", controllers.TaxController)
	}
}

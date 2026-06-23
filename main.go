package main

import (
	"cal-tex/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.Tax(r)

	// r.GET("test", test.OnTest)
	// r.POST("/tax/calculations", )

	r.Run(":5000")
}

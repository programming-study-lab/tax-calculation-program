package main

import (
	"cal-tex/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.Tax(r)

	r.Run(":5000")
}

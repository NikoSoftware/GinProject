package main

import (
	"GinProject/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.UserRoutersInit(r)

	r.Run(":7200")
}

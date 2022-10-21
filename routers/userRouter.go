package routers

import (
	"GinProject/controller"
	"GinProject/filter"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {

	router := r.Group("/user")

	router.GET("/userInfo", filter.TokenFilter, controller.UserController{}.GetUserInfo)

	router.GET("/users", filter.TokenFilter, controller.UserController{}.GetAllStudent)
}

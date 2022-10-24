package routers

import (
	"GinProject/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutersInit(r *gin.Engine) {

	router := r.Group("/user")

	router.GET("/userInfo", controller.UserController{}.GetUserInfo)

	router.GET("/users", controller.UserController{}.GetAllStudent)

	router.GET("/course", controller.UserController{}.GetCourse)

	router.GET("/teacher", controller.UserController{}.GetStudents)

	router.GET("/sc", controller.UserController{}.GetStudentCourse)

	router.GET("/redis", controller.UserController{}.GetRedis)

}

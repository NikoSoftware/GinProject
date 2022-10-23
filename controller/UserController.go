package controller

import (
	"GinProject/db"
	"GinProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, model.User{
		Name:  "小明",
		Years: 12,
		Sex:   1,
	})

}

func (u UserController) GetAllStudent(c *gin.Context) {
	var s []model.Student

	db.Db.Find(&s)

	c.JSON(http.StatusOK, s)

}

func (u UserController) GetCourse(c *gin.Context) {
	var course []model.Course

	db.Db.Debug().Model(&model.Course{}).Preload("Teacher").Find(&course)
	c.JSON(http.StatusOK, course)

}

func (u UserController) GetStudents(c *gin.Context) {
	var teachers []model.Teacher

	db.Db.Find(&teachers)
	c.JSON(http.StatusOK, teachers)

}

func (u UserController) GetStudentCourse(c *gin.Context) {
	var students []model.Student

	db.Db.Debug().Model(&model.Student{}).Preload("Sc").Preload("Sc.Course").Preload("Sc.Course.Teacher").Find(&students)

	c.JSON(http.StatusOK, students)

}

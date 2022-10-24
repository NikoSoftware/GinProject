package controller

import (
	"GinProject/db"
	"GinProject/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	b, _ := json.Marshal(students)

	db.Redisdb.Del("students")
	db.Redisdb.Set("students", b, time.Minute*50)

	c.JSON(http.StatusOK, students)

}

func (u UserController) GetRedis(c *gin.Context) {

	var students []model.Student

	studentsc := db.Redisdb.Get("students")

	fmt.Println(studentsc.Val())

	b, _ := studentsc.Bytes()

	json.Unmarshal(b, &students)

	c.JSON(http.StatusOK, students)

}

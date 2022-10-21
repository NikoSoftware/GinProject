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
	s := &model.Student{}
	db.Db.First(s)

	c.JSON(http.StatusOK, s)

}

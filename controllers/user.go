package controllers

import (
	"fmt"
	"net/http"
	"time"

	"gza/user/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db := models.GetDatabase()
	fmt.Println(db.Debug())
	r := db.Model(users).Find(&users)
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	var user models.User
	db := models.GetDatabase()
	r := db.First(&user, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	// c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}

func CreateUser(c *gin.Context) {
	var data models.User
	c.BindJSON(&data)
	data.Created = time.Now()
	data.Updated = time.Now()
	var user models.User
	db := models.GetDatabase()
	r := db.First(&user, "id = ?", data.Id)
	if r.RowsAffected != 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	db.Model(user).Create(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func Login(c *gin.Context) {
	var user models.User
	db := models.GetDatabase()
	r := db.Where(&user, "id = ?, password = ?", c.Param("id"), c.Param("password"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "아이디 또는 비밀번호가 일치하지 않습니다.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}

func UpdateUser(c *gin.Context) {
	var data models.User
	c.BindJSON(&data)
	var user models.User
	db := models.GetDatabase()
	r := db.Find(&user, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	db.Model(&user).Updates(&data)
	//models.DB.Model(&info).Update("Name", "Lee")
	//models.DB.Model(&info).Update(models.Info{Name: "Lee", Email: "rhdtha01@gmail.com"})
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	db := models.GetDatabase()
	r := db.Find(&user, "id = ?", c.Param("id"))
	if r.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "데이터가 존재하지 않습니다",
		})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}

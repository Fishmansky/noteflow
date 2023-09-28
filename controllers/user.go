package controllers

import (
	"net/http"

	"github.com/Fishmansky/noteflow/inits"
	"github.com/Fishmansky/noteflow/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	if len(body.Email) == 0 || len(body.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty login credentials",
		})
		return
	}

	user := models.User{}
	inits.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find user",
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong password",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

}

func ModifyUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}

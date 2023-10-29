package controllers

import (
	"net/http"

	"github.com/Fishmansky/noteflow/inits"
	"github.com/Fishmansky/noteflow/middleware"
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
	accessToken, err := middleware.CreateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Failed to create access token",
		})
		return
	}
	refreshToken, err := middleware.CreateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Failed to create refresh token",
		})
		return
	}
	saveAccessErr := middleware.SaveToken(user.ID, accessToken)
	if saveAccessErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Failed to save access token",
		})
		return
	}
	saveRefreshErr := middleware.SaveToken(user.ID, refreshToken)
	if saveRefreshErr != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Failed to save refresh token",
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

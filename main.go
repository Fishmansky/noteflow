package main

import (
	"github.com/Fishmansky/noteflow/controllers"
	"github.com/Fishmansky/noteflow/inits"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.ConnectToDB()
	inits.SyncDB()
	inits.ConnecRedis()
}

func main() {
	r := gin.Default()
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.GET("/register", controllers.Register)
	r.Run()
}

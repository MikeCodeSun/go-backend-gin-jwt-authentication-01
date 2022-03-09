package app

import (
	"example/server/controler"
	"example/server/util"
	
	"net/http"

	"github.com/gin-gonic/gin"
)



func protectedRoute(c *gin.Context) {

  user := c.MustGet("user")
	// greeting := fmt.Sprintf("Hello %s, welcome to website", user.name)
	c.JSON(http.StatusOK, gin.H{"msg": "This is protected!", "user": user })
}


func Routers() {
	R.POST("/register", controler.Register)
	R.POST("/login", controler.Login)
	R.GET("/protected", util.Auth(),protectedRoute)
}
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()


func getHome (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Home Page"})
}

func notFound (c *gin.Context){
c.JSON(http.StatusNotFound, gin.H{"msg": "page not found"})
}


func StartApp() {
	R.GET("/", getHome)
	Routers()
	R.NoRoute(notFound)
	R.Run(":3000")
}
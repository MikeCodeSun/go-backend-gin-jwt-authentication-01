package app

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	R.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	R.GET("/", getHome)
	Routers()
	R.NoRoute(notFound)
	R.Run(":4000")
}
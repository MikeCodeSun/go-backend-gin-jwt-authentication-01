package util

import (
	"fmt"
	"net/http"
	"os"


	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get env
		if envErr := godotenv.Load(); envErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "env wrong"})
			return
		}
	
		secret := os.Getenv("SECRET")
		

		// get cookie
		cookie, cookieErr := c.Cookie("jwt")
		if cookieErr != nil {
			fmt.Println(cookieErr)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "cookie wrong"})
			return
		}
		fmt.Println(cookie)
// get req.headers.authorization
		// authHeader := c.Request.Header["Authorization"][0]

		// fmt.Printf("header is %s", authHeader)

		// if len(strings.TrimSpace(authHeader)) == 0 || strings.HasPrefix(authHeader, "Bearer ") == false {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"msg": "token not exist"})
		// 	return
		// }  

		// token := strings.Split(authHeader, " ")[1]
		// parse token 
		ss, tErr:= jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if tErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "token wrong"})
			return
			
		} 
		//  set token info to req
		c.Set("user", ss.Claims)
		c.Next()
	}
}
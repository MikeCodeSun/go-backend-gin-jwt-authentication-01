package controler

import (
	"example/server/db"
	"example/server/util"
	"fmt"
	"time"

	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)
// register control
func Register(c *gin.Context) {
	var user db.User
	// get dot env varible
	// if envErr := godotenv.Load(); envErr != nil {
	// 	errMsg := util.CustomError("env err")
	// 	c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
	// 	return
	// }
	// secret := os.Getenv("SECRET")
	// get json data from req.body
	if bindErr := c.ShouldBindJSON(&user); bindErr != nil {
		errMsg := util.CustomError("bind err")
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}

	// bcrypt password
	hashPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if hashErr != nil {
		errMsg := util.CustomError("hash err")
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	user.Password = string(hashPassword)
	// prepare db
	stmt, preErr := db.Client.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if preErr != nil {
		errMsg := util.CustomError("pre err 11")
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	defer stmt.Close()

	
	// insert db
	res, execErr := stmt.Exec(user.Name, user.Email, user.Password)

	if execErr != nil {
		errMsg := util.CustomError("insert err")
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
//  get id
	userId, idErr := res.LastInsertId()
	if idErr != nil {
		errMsg := util.CustomError("id err")
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	user.ID = int(userId)
// res
	c.JSON(http.StatusAccepted, gin.H{"msg": "register success", "user": user })
}

// login contoler
func Login(c *gin.Context) {
	var (
		user db.User
		userDb db.User
	) 

	// get dot env varible
	if envErr := godotenv.Load(); envErr != nil {
		errMsg := util.CustomError("env err")
		fmt.Println(envErr)
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}

	secret := os.Getenv("SECRET")

	//  get log info from req body
	if jsonErr := c.ShouldBindJSON(&user); jsonErr != nil {
		errMsg := util.CustomError("json err")
		fmt.Println(errMsg)
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	//prepare db
	stmt, preErr:= db.Client.Prepare("SELECT * FROM users WHERE email=?")
	if preErr != nil {
		errMsg := util.CustomError("pre err")
		fmt.Println(preErr)
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	defer stmt.Close()
	// query row user
	ss := stmt.QueryRow(user.Email)
	if scanErr := ss.Scan(&userDb.ID, &userDb.Name, &userDb.Email ,&userDb.Password); scanErr != nil {
		errMsg := util.CustomError("scan err")
		fmt.Println(scanErr)
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	// bcrypt compare password
	if comErr := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password)); comErr != nil {
		errMsg := util.CustomError("compare password err")
		fmt.Println(comErr)
		fmt.Println(userDb.Password, user.Password)
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}

	// jwt generate token
	claims := jwt.MapClaims{}
	claims["name"] = userDb.Name
	claims["id"] = userDb.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	jwtClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign ... dotenv secret
  token, tokenErr := jwtClaim.SignedString([]byte(secret))
	if tokenErr != nil {
		errMsg := util.CustomError("token err")
		fmt.Println(tokenErr)
		c.JSON(errMsg.StatusCode, gin.H{"msg": errMsg})
		return
	}
	// set cookie
	c.SetCookie("jwt", token, 3600, "/", "localhost",false, false)

	c.JSON(http.StatusAccepted, gin.H{"msg": "login success", "user": userDb, "token": token})
}

// log out

func Logout(c *gin.Context){
	c.SetCookie("jwt", "", -1, "/", "localhost", false,false)
	c.JSON(http.StatusAccepted, gin.H{"msg": "logout"})
}
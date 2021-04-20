package main

import (
	"crypto/sha1"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var users = make(map[string]string)

func main() {
	router := gin.Default()
	router.GET("/login", GetLogin)
	router.GET("/logout", GetLogout)
	router.POST("/upload", GetUpload)
	router.GET("/status", GetStatus)
	router.Run(":8080")
}

func GetLogin(c *gin.Context) {
	var PASSWORD = "password"
	user := "user"
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 30
	random := rand.Intn(max-min+1) + min
	token := strconv.Itoa(random)

	user, ok := users[user]
	if ok == true {
		c.Abort()
	} else {
		users[user] = token

		h := sha1.New()
		h.Write([]byte(token))
		bs := h.Sum(nil)

		user, password, hasAuth := c.Request.BasicAuth()
		if password == PASSWORD && hasAuth == true {
			c.JSON(200, gin.H{
				"message": "Hi " + user + ", welcome to the DPIP System",
				"token":   bs,
			})
		} else {
			c.Abort()
		}
	}
}

func GetLogout(c *gin.Context) {
	user := "user"
	c.JSON(200, gin.H{
		"message": "Bye " + user + ", your token has been revoked",
	})
	delete(users, user)

}

func GetUpload(c *gin.Context) {
	file, _ := c.FormFile("data")
	c.SaveUploadedFile(file, "test2.jpg")
	fileName := file.Filename
	fileSize := file.Size
	c.JSON(200, gin.H{
		"message":  "An image has been successfully uploaded",
		"filename": fileName,
		"size":     fileSize,
	})
}

func GetStatus(c *gin.Context) {
	user := "user"
	t := time.Now()
	c.JSON(200, gin.H{
		"message": "Hi " + user + ", the DPIP System is Up and Running",
		"time":    t.Format("2006-01-02 3:4:5"),
	})
}

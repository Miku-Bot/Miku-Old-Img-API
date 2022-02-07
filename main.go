// Special thanks to dondish for helping me
package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	var baseURL string
	baseURL = "https://cdn.derpyenterprises.org"

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	router.GET("/takagijson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"url": baseURL + "/miku/takagi/image(" + strconv.Itoa(rand.Intn(48-1+1)+1) + ").jpeg"})
	})

	router.GET("/saojson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"url": baseURL + "/miku/sao/image(" + strconv.Itoa(rand.Intn(32-1+1)+1) + ").jpeg"})
	})

	router.GET("/ddlcjson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"url": baseURL + "/miku/ddlc/image(" + strconv.Itoa(rand.Intn(27-1+1)+1) + ").jpeg"})
	})

	router.GET("/konosubajson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"url": baseURL + "/miku/konosuba/image(" + strconv.Itoa(rand.Intn(51-1+1)+1) + ").gif"})
	})
	router.GET("/lovelivejson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"url": baseURL + "/miku/lovelive/image(" + strconv.Itoa(rand.Intn(103-1+1)+1) + ").gif"})
	})
	router.GET("/k_onjson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"url": baseURL + "/miku/k_on/image(" + strconv.Itoa(rand.Intn(142-1+1)+1) + ").gif"})
	})

	return router
}

func main() {
	router := setupRouter()

	router.Run(os.Getenv("PORT"))
}

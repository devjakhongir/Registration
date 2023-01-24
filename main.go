package main

import (
	"os"
	"log"
	"net/http"
	"app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	
	HTTP_PORT := os.Getenv("APP_HTTP_PORT")

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/change-password", controllers.PostChangePassword)
	r.GET("/echo", controllers.WebSocket)

	r.Run(HTTP_PORT)
}
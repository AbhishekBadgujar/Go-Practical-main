package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Middlewares -  if we call it in function for API call, it will set these details for that API
func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.Status(http.StatusOK)
	}

}

func ClientRoutes() {
	router := gin.Default()
	router.GET("/person/:person_id/info", getPerson)
	router.POST("/person/create", createPerson)
	if err := router.Run(":" + os.Getenv("PORT")); err != nil { // main gin engine
		log.Printf("Failed to run server:%v", err)
	}
}

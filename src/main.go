package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")

	log.Printf("This is bot token %s \n", slackBotToken)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/slack/interactivity", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Interactivity data accepted",
		})
	})

	r.POST("/slack/commands", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Commands data accepted",
		})
	})

	r.Run()

}

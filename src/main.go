package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
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

	type EventVerification struct {
		Token     string `json:"token"`
		Challenge string `json:"challenge"`
		Type      string `json:"type"`
	}

	r.POST("/slack/events", func(ctx *gin.Context) {
		jsonData, err := ctx.GetRawData()

		if err != nil {
			log.Fatalln("Error getting data from body")
		}

		var eventVerificationObject EventVerification

		json.Unmarshal(jsonData, &eventVerificationObject)

		ctx.String(200, eventVerificationObject.Challenge)
	})

	api := slack.New(slackBotToken)

	api.PostMessage("CSBJY2Z47", slack.MsgOptionText("Some Text", false), slack.MsgOptionAttachments(slack.Attachment{
		Text: "This is some text",
	}))

	r.Run()

}

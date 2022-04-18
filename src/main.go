package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackSigningSecret := os.Getenv("SLACK_SIGNING_SECRET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	api := slack.New(slackBotToken)

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

		sv, err := slack.NewSecretsVerifier(ctx.Request.Header, slackSigningSecret)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}

		if _, err := sv.Write(jsonData); err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			ctx.AbortWithStatus(500)
			return
		}

		if err := sv.Ensure(); err != nil {
			ctx.AbortWithStatusJSON(401, err)
			return
		}
		eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(jsonData), slackevents.OptionNoVerifyToken())
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}

		if eventsAPIEvent.Type == slackevents.URLVerification {
			eventData := make(map[string]interface{})

			json.Unmarshal(jsonData, &eventData)

		}

		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			api.PostMessage("CSBJY2Z47", slack.MsgOptionText("Some Text", false), slack.MsgOptionAttachments(slack.Attachment{
				Text: "This is some text",
			}))
		}
	})

	portNumber := fmt.Sprintf(":%s", port)

	r.Run(portNumber)

}

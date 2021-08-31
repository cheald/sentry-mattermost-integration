package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsalex/sentry-mattermost-sidecar/internal"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("sms")

	viper.BindEnv("mattermost_webhook_url")
	viper.BindEnv("host")
	viper.BindEnv("port")

	viper.SetDefault("addr", "0.0.0.0")
	viper.SetDefault("port", "1323")

	if viper.GetString("mattermost_webhook_url") == "" {
		log.Fatalf("SMS_MATTERMOST_WEBHOOK_URL environment variable must be set!")
	}
}

func main() {
	r := gin.Default()

	r.POST("/:channel", func(c *gin.Context) {
		channel := c.Param("channel")
		body := internal.Webhook{}

		err := c.ShouldBindJSON(&body)
		if err != nil {
			log.Fatalf("An Error Occured during bind json: %v", err)
		}

		postBody, err := json.Marshal(map[string]interface{}{
			"channel": channel,
			"attachments": []interface{}{
				map[string]interface{}{
					"title":       body.Event.Title,
					"color":       "#FF0000",
					"author_name": "Sentry",
					"author_icon": "https://assets.stickpng.com/images/58482eedcef1014c0b5e4a76.png",
					"title_link":  body.URL,
					"fields": []interface{}{
						map[string]interface{}{
							"short": false,
							"title": "Culprit",
							"value": body.Culprit,
						},
						map[string]interface{}{
							"short": false,
							"title": "Project",
							"value": body.ProjectSlug,
						},
					},
				},
			},
		})
		if err != nil {
			log.Fatalf("Error during json marshal: %v", err)
		}

		resp, err := http.Post(
			viper.GetString("mattermost_webhook_url"),
			"application/json",
			bytes.NewBuffer(postBody),
		)
		if err != nil {
			log.Fatalf("Error when performing webhook call: %v", err)
		}
		defer resp.Body.Close()
	})

	r.Run(fmt.Sprintf(
		"%s:%s",
		viper.GetString("host"),
		viper.GetString("port"),
	))
}
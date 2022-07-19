package main

import (
	"encoding/json"
	"fmt"
)

func ConvertEvent(channel string, in []byte) (out *WebhookPayload, err error) {
	var payload EventAlert
	err = json.Unmarshal(in, &payload)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling payload: %v", err)
	}
	event := payload.Data.Event

	fields := []Field{
		{
			Short: true,
			Title: "Culprit",
			Value: event.Culprit,
		},
		{
			Short: true,
			Title: "Project",
			Value: fmt.Sprintf("%d", event.Project),
		},
	}

	for _, tag := range event.Tags {
		if tag[0] != "level" {
			fields = append(fields, Field{
				Short: true,
				Title: tag[0],
				Value: tag[1],
			})
		}
	}

	fields = append(fields)

	body := &WebhookPayload{
		Channel: channel,
		Attachments: []Attachment{
			{
				Title:      event.Title,
				Color:      "#FF0000",
				AuthorName: "Sentry",
				AuthorIcon: "https://assets.stickpng.com/images/58482eedcef1014c0b5e4a76.png",
				TitleLink:  event.URL,
				Fields:     fields,
			},
		},
	}

	return body, nil
}

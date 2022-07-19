package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ProcessEvent(channel string, in []byte) (out *WebhookPayload, err error) {
	var payload Alert
	err = json.Unmarshal(in, &payload)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling payload: %v", err)
	}

	var result *WebhookPayload

	if payload.Action == "triggered" {
		result, err = TriggeredEvent(in)
	} else {
		return nil, fmt.Errorf("Unknown event: %s", payload.Action)
	}
	result.Channel = channel
	return result, err
}

func CreatedEvent(in []byte) (out *WebhookPayload, err error) {
	var payload CreatedAlert
	err = json.Unmarshal(in, &payload)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling payload: %v", err)
	}
	issue := payload.Data.Issue

	body := &WebhookPayload{
		Attachments: []Attachment{
			{
				Title:      issue.Title,
				Color:      "#FF0000",
				AuthorName: "Sentry",
				AuthorIcon: "https://assets.stickpng.com/images/58482eedcef1014c0b5e4a76.png",
				TitleLink:  fmt.Sprintf("%s", issue.Permalink),
			},
		},
	}

	return body, nil
}

func tagsToFields(tags [][]string, suppressedTags map[string]bool) []*Field {
	fields := make([]*Field, 0)
	for _, tag := range tags {
		_, exist := suppressedTags[tag[0]]
		if !exist {
			fields = append(fields, &Field{Short: true, Title: tag[0], Value: tag[1]})
		}
	}
	return fields
}

var SUPPRESSED_TAGS = map[string]bool{"os": true, "os.name": true, "level": true, "logger": true, "runtime.name": true}

func TriggeredEvent(in []byte) (out *WebhookPayload, err error) {

	var payload TriggeredAlert
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
			Value: strings.Split(event.URL, "/")[7],
		},
	}

	dynamicFields := tagsToFields(event.Tags, SUPPRESSED_TAGS)
	for _, field := range dynamicFields {
		fields = append(fields, *field)

	}

	body := &WebhookPayload{
		Attachments: []Attachment{
			{
				Title:      event.Title,
				Color:      "#FF0000",
				AuthorName: "Sentry",
				AuthorIcon: "https://assets.stickpng.com/images/58482eedcef1014c0b5e4a76.png",
				TitleLink:  event.WebURL,
				Fields:     fields,
			},
		},
	}

	return body, nil
}

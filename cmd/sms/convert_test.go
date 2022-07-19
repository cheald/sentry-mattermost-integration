package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestTriggeredEvent(t *testing.T) {
	buf, _ := os.ReadFile("test/triggered_sample.json")
	out, _ := ProcessEvent("test", buf)
	json, _ := json.MarshalIndent(out, "", "  ")
	log.Printf("%s", json)
}

func TestCreatedEvent(t *testing.T) {
	buf, _ := os.ReadFile("test/created_sample.json")
	out, _ := ProcessEvent("test", buf)
	json, _ := json.MarshalIndent(out, "", "  ")
	log.Printf("%s", json)
}

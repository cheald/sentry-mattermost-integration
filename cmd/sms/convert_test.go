package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestConvertEvent(t *testing.T) {
	buf, _ := os.ReadFile("triggered_sample.json")
	out, _ := ConvertEvent("test", buf)
	json, _ := json.MarshalIndent(out, "", "  ")
	log.Printf("%s", json)
}

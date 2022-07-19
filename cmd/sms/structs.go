package main

import "time"

// EventAlert represents an event payload
type EventAlert struct {
	Action string `json:"action"`
	Data   struct {
		Event struct {
			EventID     string      `json:"event_id"`
			Project     int         `json:"project"`
			ProjectSlug int         `json:"project_slug"`
			Release     interface{} `json:"release"`
			Dist        interface{} `json:"dist"`
			Platform    string      `json:"platform"`
			Message     string      `json:"message"`
			Datetime    time.Time   `json:"datetime"`
			Tags        [][]string  `json:"tags"`
			Metrics     struct {
				BytesIngestedEvent int `json:"bytes.ingested.event"`
				BytesStoredEvent   int `json:"bytes.stored.event"`
			} `json:"_metrics"`
			RelayProcessed bool   `json:"_relay_processed"`
			Culprit        string `json:"culprit"`
			Environment    string `json:"environment"`
			Exception      struct {
				Values []struct {
					Stacktrace struct {
						Frames []struct {
							Function        string      `json:"function"`
							AbsPath         string      `json:"abs_path"`
							Errors          interface{} `json:"errors"`
							PreContext      []string    `json:"pre_context"`
							Vars            interface{} `json:"vars"`
							Package         interface{} `json:"package"`
							ContextLine     string      `json:"context_line"`
							Symbol          interface{} `json:"symbol"`
							ImageAddr       interface{} `json:"image_addr"`
							Module          interface{} `json:"module"`
							InApp           bool        `json:"in_app"`
							SymbolAddr      interface{} `json:"symbol_addr"`
							Filename        string      `json:"filename"`
							PostContext     []string    `json:"post_context"`
							Colno           interface{} `json:"colno"`
							RawFunction     interface{} `json:"raw_function"`
							Trust           interface{} `json:"trust"`
							Data            interface{} `json:"data"`
							Platform        interface{} `json:"platform"`
							InstructionAddr interface{} `json:"instruction_addr"`
							Lineno          int         `json:"lineno"`
						} `json:"frames"`
					} `json:"stacktrace"`
					Type   string `json:"type"`
					Module string `json:"module"`
					Value  string `json:"value"`
				} `json:"values"`
			} `json:"exception"`
			Extra struct {
				Server struct {
					Runtime struct {
						Version string `json:"version"`
						Name    string `json:"name"`
					} `json:"runtime"`
					Os struct {
						KernelVersion string `json:"kernel_version"`
						Version       string `json:"version"`
						Build         string `json:"build"`
						Name          string `json:"name"`
					} `json:"os"`
				} `json:"server"`
			} `json:"extra"`
			Fingerprint    []string `json:"fingerprint"`
			GroupingConfig struct {
				Enhancements string `json:"enhancements"`
				ID           string `json:"id"`
			} `json:"grouping_config"`
			Hashes   []string `json:"hashes"`
			KeyID    string   `json:"key_id"`
			Level    string   `json:"level"`
			Location string   `json:"location"`
			Logger   string   `json:"logger"`
			Metadata struct {
				Function string `json:"function"`
				Type     string `json:"type"`
				Value    string `json:"value"`
				Filename string `json:"filename"`
			} `json:"metadata"`
			Received float64 `json:"received"`
			Sdk      struct {
				Version      string      `json:"version"`
				Name         string      `json:"name"`
				Packages     interface{} `json:"packages"`
				Integrations interface{} `json:"integrations"`
			} `json:"sdk"`
			Timestamp float64 `json:"timestamp"`
			Title     string  `json:"title"`
			Type      string  `json:"type"`
			Version   string  `json:"version"`
			URL       string  `json:"url"`
			WebURL    string  `json:"web_url"`
			IssueURL  string  `json:"issue_url"`
		} `json:"event"`
		TriggeredRule string `json:"triggered_rule"`
	} `json:"data"`
	Installation struct {
		UUID string `json:"uuid"`
	} `json:"installation"`
	Actor struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"actor"`
}

type WebhookPayload struct {
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Color      string  `json:"color"`
	AuthorName string  `json:"author_name"`
	AuthorIcon string  `json:"author_icon"`
	Fields     []Field `json:"fields"`
}

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

package main

import "time"

type Alert struct {
	Action string `json:"action"`
}

// TriggeredAlert represents an event payload for a triggered issue
type TriggeredAlert struct {
	Action       string `json:"action"`
	Installation struct {
		UUID string `json:"uuid"`
	} `json:"installation"`
	Data struct {
		Event struct {
			EventID  string      `json:"event_id"`
			Project  int         `json:"project"`
			Release  interface{} `json:"release"`
			Dist     interface{} `json:"dist"`
			Platform string      `json:"platform"`
			Message  string      `json:"message"`
			Datetime time.Time   `json:"datetime"`
			Tags     [][]string  `json:"tags"`
			Metrics  struct {
				BytesIngestedEvent int `json:"bytes.ingested.event"`
				BytesStoredEvent   int `json:"bytes.stored.event"`
			} `json:"_metrics"`
			Ref        int `json:"_ref"`
			RefVersion int `json:"_ref_version"`
			Contexts   struct {
				Os struct {
					Name          string `json:"name"`
					Version       string `json:"version"`
					Build         string `json:"build"`
					KernelVersion string `json:"kernel_version"`
					Type          string `json:"type"`
				} `json:"os"`
				Runtime struct {
					Name    string `json:"name"`
					Version string `json:"version"`
					Type    string `json:"type"`
				} `json:"runtime"`
			} `json:"contexts"`
			Culprit     string `json:"culprit"`
			Environment string `json:"environment"`
			Exception   struct {
				Values []struct {
					Type       string `json:"type"`
					Value      string `json:"value"`
					Module     string `json:"module"`
					Stacktrace struct {
						Frames []struct {
							Function string `json:"function"`
							Filename string `json:"filename"`
							AbsPath  string `json:"abs_path"`
							Lineno   int    `json:"lineno"`
							InApp    bool   `json:"in_app"`
						} `json:"frames"`
					} `json:"stacktrace"`
					ThreadID int `json:"thread_id"`
				} `json:"values"`
			} `json:"exception"`
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
				DisplayTitleWithTreeLabel bool   `json:"display_title_with_tree_label"`
				Filename                  string `json:"filename"`
				Function                  string `json:"function"`
				Type                      string `json:"type"`
				Value                     string `json:"value"`
			} `json:"metadata"`
			Modules struct {
				ZxcvbnRuby string `json:"zxcvbn-ruby"`
			} `json:"modules"`
			NodestoreInsert float64 `json:"nodestore_insert"`
			Received        float64 `json:"received"`
			Sdk             struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"sdk"`
			Threads struct {
				Values []struct {
					ID      int  `json:"id"`
					Crashed bool `json:"crashed"`
					Current bool `json:"current"`
				} `json:"values"`
			} `json:"threads"`
			Timestamp float64 `json:"timestamp"`
			Title     string  `json:"title"`
			Type      string  `json:"type"`
			Version   string  `json:"version"`
			URL       string  `json:"url"`
			WebURL    string  `json:"web_url"`
			IssueURL  string  `json:"issue_url"`
			IssueID   string  `json:"issue_id"`
		} `json:"event"`
		TriggeredRule string `json:"triggered_rule"`
	} `json:"data"`
	Actor struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"actor"`
}

// CreatedAlert represents an event payload for a new issue creation event
type CreatedAlert struct {
	Action       string `json:"action"`
	Installation struct {
		UUID string `json:"uuid"`
	} `json:"installation"`
	Data struct {
		Issue struct {
			ID            string      `json:"id"`
			ShareID       interface{} `json:"shareId"`
			ShortID       string      `json:"shortId"`
			Title         string      `json:"title"`
			Culprit       string      `json:"culprit"`
			Permalink     interface{} `json:"permalink"`
			Logger        interface{} `json:"logger"`
			Level         string      `json:"level"`
			Status        string      `json:"status"`
			StatusDetails struct {
			} `json:"statusDetails"`
			IsPublic bool   `json:"isPublic"`
			Platform string `json:"platform"`
			Project  struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
				Platform string `json:"platform"`
			} `json:"project"`
			Type     string `json:"type"`
			Metadata struct {
				Value                     string `json:"value"`
				Type                      string `json:"type"`
				Filename                  string `json:"filename"`
				Function                  string `json:"function"`
				DisplayTitleWithTreeLabel bool   `json:"display_title_with_tree_label"`
			} `json:"metadata"`
			NumComments         int           `json:"numComments"`
			AssignedTo          interface{}   `json:"assignedTo"`
			IsBookmarked        bool          `json:"isBookmarked"`
			IsSubscribed        bool          `json:"isSubscribed"`
			SubscriptionDetails interface{}   `json:"subscriptionDetails"`
			HasSeen             bool          `json:"hasSeen"`
			Annotations         []interface{} `json:"annotations"`
			IsUnhandled         bool          `json:"isUnhandled"`
			Count               string        `json:"count"`
			UserCount           int           `json:"userCount"`
			FirstSeen           time.Time     `json:"firstSeen"`
			LastSeen            time.Time     `json:"lastSeen"`
		} `json:"issue"`
	} `json:"data"`
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

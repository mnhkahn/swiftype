package swiftype

// Swiftype
type SwiftypeResult struct {
	Errors struct{} `json:"errors"`
	Info   struct {
		Page struct {
			CurrentPage      int      `json:"current_page"`
			Facets           struct{} `json:"facets"`
			NumPages         int      `json:"num_pages"`
			PerPage          int      `json:"per_page"`
			Query            string   `json:"query"`
			TotalResultCount int      `json:"total_result_count"`
		} `json:"page"`
	} `json:"info"`
	RecordCount int `json:"record_count"`
	Records     struct {
		Page []struct {
			Explanation interface{} `json:"_explanation"`
			Index       string      `json:"_index"`
			Score       float64     `json:"_score"`
			Typee       string      `json:"_type"`
			Version     interface{} `json:"_version"`
			Body        string      `json:"body"`
			ExternalID  string      `json:"external_id"`
			Highlight   struct {
				Body  string `json:"body"`
				Title string `json:"title"`
			} `json:"highlight"`
			ID          string        `json:"id"`
			Image       string        `json:"image"`
			Info        string        `json:"info"`
			Popularity  int           `json:"popularity"`
			PublishedAt string        `json:"published_at"`
			Sections    []interface{} `json:"sections"`
			Sort        interface{}   `json:"sort"`
			Title       string        `json:"title"`
			Type        string        `json:"type"`
			UpdatedAt   string        `json:"updated_at"`
			URL         string        `json:"url"`
		} `json:"page"`
	} `json:"records"`
}

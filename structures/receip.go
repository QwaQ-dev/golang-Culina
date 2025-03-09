package structures

type Recipes struct {
	Id           int               `json:"id"`
	Name         string            `json:"name"`
	Descr        string            `json:"descr"`
	Diff         string            `json:"diff"` //difficult
	Filters      []string          `json:"filters"`
	Imgs         map[string]string `json:"imgs"`
	AuthorID     int               `json:"authorID"`
	Ingredients  map[string]string `json:"ingredients"`
	Steps        map[string]string `json:"steps"`
	Review_count int               `json:"review_count, omitempty"`
	Avg_rating   float64           `json:"avg_rating, omitempty"`
	Reviews      []Review          `json:"reviews, omitempty"`
	Created_at   string            `json:"created_at, omitempty"`
}

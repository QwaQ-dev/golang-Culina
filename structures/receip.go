package structures

type Recipes struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Descr       string            `json:"descr"`
	Diff        string            `json:"diff"` //difficult
	Filters     []string          `json:"filters"`
	Imgs        map[string]string `json:"imgs"`
	AuthorID    int               `json:"authorID"`
	Ingredients map[string]string `json:"ingredients"`
	Steps       map[string]string `json:"steps"`
}

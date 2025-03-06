package structures

type Recipes struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Descr       string            `json:"descr"`
	Diff        string            `json:"diff"`
	Filters     map[string]string `json:"filters"`
	Imgs        map[int]string    `json:"imgs"`
	Author      string            `json:"author"`
	Ingredients map[string]string `json:"ingredients"`
	Steps       map[string]string `json:"steps"`
}

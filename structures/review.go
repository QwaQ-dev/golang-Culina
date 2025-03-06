package structures

type Review struct {
	Id           int    `json:"id"`
	Review_text  string `json:"review_text"`
	Rating_value int    `json:"rating_value"`
	Reviewed_by  string `json:"reviewed_by"`
}

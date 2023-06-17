package models

type POSTS struct {
	id         int     `json:"id"`
	desc       string  `json:"desc"`
	title      string  `json:"titme"`
	img        string  `json:"img"`
	updated_at float32 `json:"updated_at"`
}

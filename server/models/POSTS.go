package models

type POSTS struct {
	Id         int     `json:"id"`
	Desc       string  `json:"desc"`
	Title      string  `json:"title"`
	Img        *string `json:"img,omitempty"`
	Updated_at float32 `json:"updated_at"`
}

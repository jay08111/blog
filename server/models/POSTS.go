package models

type POSTS struct {
	Id         int     `json:"id"`
	Desc       string  `json:"description"`
	Title      string  `json:"title"`
	Img        *string `json:"img,omitempty"`
	Updated_At float32 `json:"updated_at"`
}

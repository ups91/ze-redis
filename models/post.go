package models

const T = "post"

type Post struct {
	PostName string `json:"post_name"`
	Author   string `json:"author"`
	Date     string `json:"date"`
}

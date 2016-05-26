package models

type news struct {
	Id int `json:"id"`
	CreatedAt float32 `json:"createdAt"`
	Title string `json:"title"`
	Body string `json:"body"`
}
